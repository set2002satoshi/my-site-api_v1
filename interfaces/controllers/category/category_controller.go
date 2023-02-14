package category

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	DBCategory "github.com/set2002satoshi/my-site-api/interfaces/database/category"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/my-site-api/usecase/category"
)

type CategoryController struct {
	Interactor usecase.CategoryInteractor
}

func NewCategoryController(db database.DB) *CategoryController {
	return &CategoryController{
		Interactor: usecase.CategoryInteractor{
			DB:           &database.DBRepository{DB: db},
			CategoryRepo: &DBCategory.CategoryRepository{},
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
