//go:generate go tool mockgen -source=$GOFILE -package=$GOPACKAGE_test -destination=./mocks/$GOFILE

package repository

import (
	"context"

	"go-clean-grpc/internal/domain/model"
)

// ユーザリポジトリ
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
