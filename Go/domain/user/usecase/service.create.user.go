package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
)

type CreateUserRequest struct {
	Phone string			`json:"phone"`
	Name	string			`json:"name"`
	Role	string			`json:"role"`
}

func (s *Service) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {

	registeredUser, err := s.users.FindUserByPhone(ctx, user.Phone)
	if err != nil {
		return nil, err
	}
	if registeredUser != nil {
		return nil, errors.New("duplicate_user")
	}
	fmt.Printf("%v", registeredUser)

	newUser, err := s.users.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return newUser, nil
}