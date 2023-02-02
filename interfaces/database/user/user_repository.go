package user

import (
	"errors"

	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error) {
	if result := db.Create(obj); result.Error != nil {
		return &models.UserEntity{}, errors.New("create user failed")
	}
	// userが作成されたか確認
	var createdUser *models.UserEntity
	result := db.Where("user_id = ?", obj.GetUserId()).First(&createdUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &models.UserEntity{}, errors.New("created user but user found")
	}
	return createdUser, nil
}
