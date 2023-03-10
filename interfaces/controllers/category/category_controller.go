package category

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	"github.com/set2002satoshi/my-site-api/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api/usecase/service"
)

type CategoryController struct {
	Interactor service.CategoryInteractor
}

func NewCategoryController(db config.DB) *CategoryController {
	return &CategoryController{
		Interactor: service.CategoryInteractor{
			DB:           &config.DBRepository{DB: db},
			CategoryRepo: &database.CategoryRepository{},
		},
	}
}

func (bc *CategoryController) convertActiveCategoryToDTO(obj *models.CategoryEntity) response.ActiveCategoryEntity {
	return response.ActiveCategoryEntity{
		Id:           int(obj.GetCategoryID()),
		CategoryName: obj.GetCategoryName(),
	}
}

func (bc *CategoryController) convertActiveCategoryToDTOs(obj []*models.CategoryEntity) []response.ActiveCategoryEntity {
	CEL := make([]response.ActiveCategoryEntity, len(obj))
	for i, c := range obj {
		category := response.ActiveCategoryEntity{
			Id:           int(c.GetCategoryID()),
			CategoryName: c.GetCategoryName(),
		}
		CEL[i] = category
	}
	return CEL

}
