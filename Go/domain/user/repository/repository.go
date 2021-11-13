package repository

import (
	"context"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
}
