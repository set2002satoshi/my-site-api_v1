package blog

import (
	"github.com/set2002satoshi/my-site-api/models"
	"gorm.io/gorm"
)


type BlogRepository interface {
	Create(db *gorm.DB, obj *models.BlogEntity) (*models.BlogEntity, error)
	GetById(db *gorm.DB, id int) (models.BlogEntity, error)
	GetAll(db *gorm.DB) ([]*models.BlogEntity, error)
	Delete(db *gorm.DB, id int) error
}
