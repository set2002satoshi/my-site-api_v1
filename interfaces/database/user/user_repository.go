package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error) {
	if result := db.Create(obj); result.Error != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0002, gorm.ErrInvalidData.Error())
	}
	// userが作成されたか確認
	var createdUser *models.UserEntity
	result := db.Where("user_id = ?", obj.GetUserId()).First(&createdUser)
	if result.Error != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0002, gorm.ErrRecordNotFound.Error())
	}
	return createdUser, nil
}
