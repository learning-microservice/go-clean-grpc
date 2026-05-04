package user_test

import (
	"context"
	stderrors "errors"
	"testing"

	"go-clean-grpc/internal/domain/errors"
	"go-clean-grpc/internal/domain/model"
	mock_repository "go-clean-grpc/internal/domain/repository/mocks"
	"go-clean-grpc/internal/usecase/user"
	mock_user "go-clean-grpc/internal/usecase/user/mocks"
	"go-clean-grpc/pkg/bcript"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestLoginUsecase_Execute(t *testing.T) {
	t.Parallel()

	const (
		id       = "U001"
		email    = "test@example.com"
		password = "password123"
		token    = "TOKEN-XXX-XXX"
	)

	passwordHash, err := bcript.ToHash(password)
	require.NoError(t, err)

	type mocks struct {
		userRepo *mock_repository.MockUserRepository
		tokenGen *mock_user.MockTokenGenerator
	}

	tests := []struct {
		name      string
		setupMock func(mocks *mocks)
		input     *user.LoginInput
		want      *user.LoginOutput
		wantErr   error
	}{
		{
			name: "success",
			setupMock: func(mocks *mocks) {
				mocks.userRepo.EXPECT().
					FindByEmail(gomock.Any(), email).
					Return(model.NewUser(id, "testman", email, passwordHash), nil)

				mocks.tokenGen.EXPECT().
					Generate(id).
					Return(token, nil)
			},
			input: &user.LoginInput{
				Email:         email,
				PlainPassword: password,
			},
			want: &user.LoginOutput{
				Token: token,
			},
		},
		{
			name: "user not found returns invalid credentials",
			setupMock: func(mocks *mocks) {
				mocks.userRepo.EXPECT().
					FindByEmail(gomock.Any(), email).
					Return(nil, stderrors.New("not found"))
			},
			input: &user.LoginInput{
				Email:         email,
				PlainPassword: password,
			},
			wantErr: errors.ErrInvalidCredentials,
		},
		{
			name: "wrong password returns invalid credentials",
			setupMock: func(mocks *mocks) {
				mocks.userRepo.EXPECT().
					FindByEmail(gomock.Any(), email).
					Return(model.NewUser(id, "testman", email, passwordHash), nil)
			},
			input: &user.LoginInput{
				Email:         email,
				PlainPassword: "wrong-password",
			},
			wantErr: errors.ErrInvalidCredentials,
		},
		{
			name: "jwt generate failure returns unexpected error",
			setupMock: func(mocks *mocks) {
				mocks.userRepo.EXPECT().
					FindByEmail(gomock.Any(), email).
					Return(model.NewUser(id, "testman", email, passwordHash), nil)

				mocks.tokenGen.EXPECT().
					Generate(id).
					Return("", stderrors.New("jwt failure"))
			},
			input: &user.LoginInput{
				Email:         email,
				PlainPassword: password,
			},
			wantErr: errors.ErrUnexpectedError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocks := mocks{
				userRepo: mock_repository.NewMockUserRepository(ctrl),
				tokenGen: mock_user.NewMockTokenGenerator(ctrl),
			}

			tt.setupMock(&mocks)

			uc := user.NewLoginUsecase(mocks.userRepo, mocks.tokenGen)
			got, gotErr := uc.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				require.Error(t, gotErr)
				require.ErrorIs(t, gotErr, tt.wantErr)
				require.Nil(t, got)
				return
			}

			require.NoError(t, gotErr)
			require.NotNil(t, got)
			require.Equal(t, tt.want.Token, got.Token)
		})
	}
}
