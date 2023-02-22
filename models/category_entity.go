package models

import (
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
)

type CategoryEntity struct {
	CategoryID   types.IDENTIFICATION    `gorm:"primaryKey"`
	CategoryName string                  `gorm:"unique;not null;max:18"`
	BlogIds      []BlogAndCategoryEntity `gorm:"foreignKey:CategoryId"`
	blogs        []BlogEntity            `gorm:"-:migration"`
}

func NewCategoryEntity(
	categoryId int,
	categoryName string,
	blogs []BlogEntity,
) (*CategoryEntity, error) {
	ce := new(CategoryEntity)
	var err error
	err = errors.Combine(err, ce.setCategoryID(categoryId))
	err = errors.Combine(err, ce.setCategoryName(categoryName))
	err = errors.Combine(err, ce.setBlogs(blogs))
	if err != nil {
		return new(CategoryEntity), err
	}
	return ce, nil
}

func (ce *CategoryEntity) GetCategoryID() types.IDENTIFICATION {
	return ce.CategoryID
}

func (ce *CategoryEntity) GetCategoryName() string {
	return ce.CategoryName
}

func (ce *CategoryEntity) GetBlogs() []BlogEntity {
	return ce.blogs
}

func (ce *CategoryEntity) setCategoryID(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	ce.CategoryID = i
	return nil
}

func (ce *CategoryEntity) setCategoryName(CategoryName string) error {
	ce.CategoryName = CategoryName
	return nil
}
func (ce *CategoryEntity) setBlogs(blogs []BlogEntity) error {
	ce.blogs = blogs
	return nil
}
