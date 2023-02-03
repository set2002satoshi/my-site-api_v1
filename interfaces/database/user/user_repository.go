package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) GetAll(db *gorm.DB) ([]*models.UserEntity, error) {
	var user []*models.UserEntity
	result := db.Find(&user)
	if result.Error != nil {
		return user, errors.Wrap(errors.NewCustomError(), errors.REPO0004, gorm.ErrRecordNotFound.Error())
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error) {
	createResult := db.Create(obj)
	if createResult.Error != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0001, gorm.ErrInvalidData.Error())
	}
	// userが作成されたか確認
	var createdUser *models.UserEntity
	result := db.Where("user_id = ?", obj.GetUserId()).First(&createdUser)
	if result.Error != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0002, gorm.ErrRecordNotFound.Error())
	}
	return createdUser, nil
}

func (repo *UserRepository) FetchEmailNumber(db *gorm.DB, email string) (int64, error) {
	var count int64
	result := db.Model(&models.UserEntity{}).Where("email = ?", email).Count(&count)
	return count, result.Error
}
