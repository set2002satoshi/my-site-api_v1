package category

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetById(db *gorm.DB, id int) (*models.CategoryEntity, error)
	GetAll(db *gorm.DB) ([]*models.CategoryEntity, error)
	Create(db *gorm.DB, obj *models.CategoryEntity) (*models.CategoryEntity, error)
}
