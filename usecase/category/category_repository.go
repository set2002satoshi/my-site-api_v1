package category

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(db *gorm.DB, obj *models.CategoryEntity) (*models.CategoryEntity, error)
}
