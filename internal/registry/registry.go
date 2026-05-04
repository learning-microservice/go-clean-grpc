package registry

import (
	"go-clean-grpc/config"
	"go-clean-grpc/internal/infra/memory"
	"go-clean-grpc/internal/usecase/user"
	"go-clean-grpc/pkg/jwt"
)

func New(cfg *config.Config) (*Registry, error) {
	// setup jwt
	jwtMng, err := jwt.New(cfg.JWT.Secret,
		jwt.Duration(cfg.JWT.Duration),
	)
	if err != nil {
		return nil, err
	}

	// setup repositories
	userRepo := memory.NewUserRepository()

	registry := Registry{}
	registry.UsecaseSet.UserLogin = user.NewLoginUsecase(userRepo, jwtMng)

	return &registry, nil
}

type Registry struct {
	UsecaseSet struct {
		UserLogin *user.LoginUsecase
	}
}
