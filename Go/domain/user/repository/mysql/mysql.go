package mysql

import (
	"github.com/kevintanuhardi/efishery_backend_test/domain/user/repository"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB // we can change all *gorm.DB to *sqlx.DB, it's optional.
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &repo{db}
}
