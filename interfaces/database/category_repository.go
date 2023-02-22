package database

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type CategoryRepository struct{}

func (repo *CategoryRepository) GetAll(db *gorm.DB) ([]*models.CategoryEntity, error) {
	var categories []*models.CategoryEntity
	if err := db.Find(&categories).Error; err != nil {
		return []*models.CategoryEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0014, gorm.ErrRecordNotFound.Error())
	}
	return categories, nil
}

func (repo CategoryRepository) GetById(db *gorm.DB, id int) (obj *models.CategoryEntity, err error) {
	if err := db.Where("category_id = ?", id).First(&obj).Error; err != nil {
		return &models.CategoryEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0013, gorm.ErrRecordNotFound.Error())
	}
	return obj, err
}

func (repo *CategoryRepository) Create(db *gorm.DB, obj *models.CategoryEntity) (*models.CategoryEntity, error) {
	if err := db.Create(&obj).Error; err != nil {
		return &models.CategoryEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0012, gorm.ErrInvalidData.Error())
	}
	return obj, nil
}

func (repo *CategoryRepository) DeleteById(db *gorm.DB, id int) error {
	if err := db.Unscoped().Delete(&models.CategoryEntity{}, id).Error; err != nil {
		return errors.Add(errors.NewCustomError(), errors.REPO0005)
	}
	return nil
}
