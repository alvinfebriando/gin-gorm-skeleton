package repository

import (
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[entity.User]
}

type userRepository struct {
	*baseRepository[entity.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:             db,
		baseRepository: &baseRepository[entity.User]{db: db},
	}
}
