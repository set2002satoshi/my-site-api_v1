package repository

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type BlogWithCategoryRepository interface {
	// GetAll(db *gorm.DB) ([]*models.BlogAndCategoryEntity, error)
	// GetById(db *gorm.DB, id int) (obj *models.BlogAndCategoryEntity, err error)
	Create(tx *gorm.DB, obj *models.BlogAndCategoryEntity) error
}
