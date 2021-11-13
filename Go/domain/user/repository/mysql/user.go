package mysql

import (
	"context"

	"github.com/kevintanuhardi/efishery_backend_test/domain/user/entity"
)

func (r *repo) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	result := r.db.Create(&user)

	if(result.Error != nil) {
		return nil, result.Error
	}

	return &user, nil
}

func (r *repo) FindUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	user := &entity.User{}

	r.db.Where("phone = ?", phone).First(user)
	return user, nil
}

