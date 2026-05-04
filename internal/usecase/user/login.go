//go:generate go tool mockgen -source=$GOFILE -package=$GOPACKAGE_test -destination=./mocks/$GOFILE
package user

import (
	"context"

	"go-clean-grpc/internal/domain/errors"
	"go-clean-grpc/internal/domain/repository"
)

type TokenGenerator interface {
	Generate(userID string) (string, error)
}

// NewLoginUsecase -.
func NewLoginUsecase(userRepo repository.UserRepository, jwtGen TokenGenerator) *LoginUsecase {
	return &LoginUsecase{
		userRepo: userRepo,
		jwtGen:   jwtGen,
	}
}

// LoginInput -.
type LoginInput struct {
	Email         string
	PlainPassword string
}

// LoginOutput -.
type LoginOutput struct {
	Token string
}

// LoginUsecase -.
type LoginUsecase struct {
	userRepo repository.UserRepository
	jwtGen   TokenGenerator
}

// Execute -.
func (uc *LoginUsecase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	if err := user.VerifyPassword(input.PlainPassword); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	token, err := uc.jwtGen.Generate(user.ID())
	if err != nil {
		return nil, errors.ErrUnexpectedError
	}

	return &LoginOutput{Token: token}, nil
}
