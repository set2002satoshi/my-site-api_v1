package category

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type CategoryRepository struct{}

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
