package blog

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)


type BlogRepository interface {
	Create(db *gorm.DB, obj *models.BlogEntity) (*models.BlogEntity, error)
}
