package usecase

import (
	"context"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
)

type CreateUserRequest struct {
	Phone string			`json:"phone"`
	Name	string			`json:"name"`
	Role	string			`json:"role"`
}

func (s *Service) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {

	newUser, err := s.users.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return newUser, nil
}