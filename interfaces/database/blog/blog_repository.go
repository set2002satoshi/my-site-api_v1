package blog

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"gorm.io/gorm"
)

type BlogRepository struct{}

func (repo *BlogRepository) GetById(db *gorm.DB, id int) (models.BlogEntity, error) {
	var findBlog models.BlogEntity
	if err := db.Where("blog_id = ?", id).First(&findBlog).Error; err != nil {
		return models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0007, gorm.ErrRecordNotFound.Error())
	}
	return findBlog, nil
}

func (repo *BlogRepository) Create(db *gorm.DB, obj *models.BlogEntity) (*models.BlogEntity, error) {
	if err := db.Create(obj).Error; err != nil {
		return &models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0006, gorm.ErrInvalidData.Error())
	}
	var createdBlog *models.BlogEntity
	if err := db.Where("blog_id = ?", obj.GetBlogId()).First(&createdBlog).Error; err != nil {
		return &models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0007, gorm.ErrRecordNotFound.Error())
	}
	return createdBlog, nil
}
