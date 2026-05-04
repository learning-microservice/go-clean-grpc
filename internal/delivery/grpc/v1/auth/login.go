package auth

import (
	"context"

	v1 "go-clean-grpc/api/v1"
	"go-clean-grpc/api/v1/v1connect"
	"go-clean-grpc/internal/domain/errors"
	"go-clean-grpc/internal/usecase/user"

	"connectrpc.com/connect"
)

// loginService -.
type loginService interface {
	Execute(ctx context.Context, input *user.LoginInput) (*user.LoginOutput, error)
}

func NewLoginHandler(login loginService) v1connect.AuthServiceHandler {
	return &loginHandler{
		login: login,
	}
}

type loginHandler struct {
	login loginService
}

func (h *loginHandler) Login(
	ctx context.Context,
	req *connect.Request[v1.LoginRequest],
) (*connect.Response[v1.LoginResponse], error) {
	// Call usecase function
	output, err := h.login.Execute(ctx, &user.LoginInput{
		Email:         req.Msg.Email,
		PlainPassword: req.Msg.Password,
	})

	// handle error response
	if err != nil {
		if errors.Is(err, errors.ErrInvalidCredentials) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&v1.LoginResponse{
		Token: output.Token,
	})
	return res, nil
}
