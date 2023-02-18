package models

import (
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
)

type CategoryEntity struct {
	CategoryID   types.IDENTIFICATION    `gorm:"primaryKey"`
	CategoryName string                  `gorm:"unique;not null;max:18"`
	Blogs        []BlogAndCategoryEntity `gorm:"foreignKey:CategoryId"`
}

func NewCategoryEntity(
	categoryId int,
	categoryName string,
) (*CategoryEntity, error) {
	ce := new(CategoryEntity)
	var err error
	err = errors.Combine(err, ce.setCategoryID(categoryId))
	err = errors.Combine(err, ce.setCategoryName(categoryName))
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
