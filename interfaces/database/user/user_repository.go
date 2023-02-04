package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) GetById(db *gorm.DB, id int) (*models.UserEntity, error) {
	var findUser *models.UserEntity
	if err := db.Where("user_id = ?", id).First(&findUser).Error; err != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0001, gorm.ErrRecordNotFound.Error())
	}
	return findUser, nil
}

func (repo *UserRepository) GetAll(db *gorm.DB) ([]*models.UserEntity, error) {
	var user []*models.UserEntity
	if err := db.Find(&user).Error; err != nil {
		return user, errors.Wrap(errors.NewCustomError(), errors.REPO0004, gorm.ErrRecordNotFound.Error())
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error) {
	if err := db.Create(obj).Error; err != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0001, gorm.ErrInvalidData.Error())
	}
	// userが作成されたか確認
	var createdUser *models.UserEntity
	if err := db.Where("user_id = ?", obj.GetUserId()).First(&createdUser).Error; err != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0002, gorm.ErrRecordNotFound.Error())
	}
	return createdUser, nil
}

func (repo *UserRepository) Update(tx *gorm.DB, obj *models.UserEntity) (*models.UserEntity, error) {
	if err := tx.Select("email", "user_name", "password", "roll", "revision").Updates(&obj).Error; err != nil {
		return &models.UserEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0003, err.Error())
	}
	return obj, nil
}

func (repo *UserRepository) FetchEmailNumber(db *gorm.DB, email string) (int64, error) {
	var count int64
	err := db.Model(&models.UserEntity{}).Where("email = ?", email).Count(&count).Error
	return count, err
}
