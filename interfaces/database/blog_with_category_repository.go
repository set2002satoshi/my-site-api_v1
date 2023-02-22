package database

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type BlogWithCategoryRepository struct{}

func (repo *BlogWithCategoryRepository) Create(db *gorm.DB, obj *models.BlogAndCategoryEntity) error {
	if err := db.Create(&obj).Error; err != nil {
		return errors.Wrap(errors.NewCustomError(), errors.REPO0015, gorm.ErrInvalidData.Error())
	}
	return nil
}
