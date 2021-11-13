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
