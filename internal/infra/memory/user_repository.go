package memory

import (
	"context"
	"fmt"
	"go-clean-grpc/internal/domain/errors"
	"go-clean-grpc/internal/domain/model"
	"go-clean-grpc/internal/domain/repository"
	"go-clean-grpc/pkg/bcript"
	"go-clean-grpc/pkg/syncmap"
)

//nolint:errcheck // This is a default password for memory repository
var _defaultPasswodHash, _ = bcript.ToHash("password")

type userRepository struct {
	users syncmap.Map[model.User]
}

// NewUserRepository -.
func NewUserRepository() repository.UserRepository {
	return &userRepository{
		users: *syncmap.New(100, map[string]*model.User{
			"yuki.kawamura@example.com": model.NewUser("U001", "Yuki Kawamura", "yuki.kawamura@example.com", _defaultPasswodHash),
		}),
	}
}

// FindByEmail -.
func (r *userRepository) FindByEmail(_ context.Context, email string) (*model.User, error) {
	value, ok := r.users.Load(email)
	fmt.Printf("find by email: %+v\n", value)
	if !ok {
		return nil, errors.ErrNotFound
	}
	return value, nil
}

// Save -.
func (r *userRepository) Save(_ context.Context, user *model.User) error {
	r.users.Store(user.Email(), user)
	return nil
}
