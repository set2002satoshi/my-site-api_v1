package category

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
)

type CategoryInteractor struct {
	DB           usecase.DBRepository
	CategoryRepo CategoryRepository
}

func (ci *CategoryInteractor) Register(obj *models.CategoryEntity) (*models.CategoryEntity, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.Create(db, obj)
}