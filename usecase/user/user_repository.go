package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error)
}
