package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(db *gorm.DB, id int) (*models.UserEntity, error)
	GetAll(db *gorm.DB) ([]*models.UserEntity, error)
	Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error)
	Update(tx *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error)
	FetchEmailNumber(db *gorm.DB, email string) (int64, error)
}
