package database

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

func (repo *BlogRepository) GetAll(db *gorm.DB) (blog []*models.BlogEntity, err error) {
	if err = db.Find(&blog).Error; err != nil {
		err = errors.Wrap(errors.NewCustomError(), errors.REPO0010, gorm.ErrRecordNotFound.Error())
		return blog, err
	}
	return
}

func (repo *BlogRepository) Delete(db *gorm.DB, id int) error {
	if err := db.Unscoped().Delete(&models.BlogEntity{}, id).Error; err != nil {
		return errors.Add(errors.NewCustomError(), errors.REPO0005)
	}
	return nil
}

func (repo *BlogRepository) Update(tx *gorm.DB, obj *models.BlogEntity) (*models.BlogEntity, error) {
	if err := tx.Select("title", "content", "revision", "updatedAt").Updates(&obj).Error; err != nil {
		return &models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0008, err.Error())
	}
	return obj, nil
}

func (repo *BlogRepository) Create(db *gorm.DB, obj *models.BlogEntity) (*models.BlogEntity, error) {
	if err := db.Create(&obj).Error; err != nil {
		return &models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0006, gorm.ErrInvalidData.Error())
	}
	var createdBlog *models.BlogEntity
	if err := db.Where("blog_id = ?", obj.GetBlogId()).First(&createdBlog).Error; err != nil {
		return &models.BlogEntity{}, errors.Wrap(errors.NewCustomError(), errors.REPO0007, gorm.ErrRecordNotFound.Error())
	}
	return createdBlog, nil
}
