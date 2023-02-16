package category

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
)

type CategoryInteractor struct {
	DB           usecase.DBRepository
	CategoryRepo CategoryRepository
}

func (ci *CategoryInteractor) FindById(id int) (*models.CategoryEntity, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.GetById(db, id)
}

func (ci *CategoryInteractor) FindAll() ([]*models.CategoryEntity, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.GetAll(db)
}

func (ci *CategoryInteractor) Register(obj *models.CategoryEntity) (*models.CategoryEntity, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.Create(db, obj)
}

func (ci *CategoryInteractor) Delete(id int) error {
	db := ci.DB.Connect()
	return ci.CategoryRepo.DeleteById(db, id)
}