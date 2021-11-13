package usecase

import (
	"context"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/repository"
)

type Service struct {
	users repository.Repository
}
type ServiceManager interface {
	createUser(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, id int) (*entity.User, error)
}

func NewService(users repository.Repository) *Service {
	return &Service{users}
}
