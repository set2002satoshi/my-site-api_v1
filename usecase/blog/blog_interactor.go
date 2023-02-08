package blog

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
)

type BlogInteractor struct {
	DB       usecase.DBRepository
	BlogRepo BlogRepository
}

func (bi *BlogInteractor) Register(obj *models.BlogEntity) (*models.BlogEntity, error) {
	db := bi.DB.Connect()
	return bi.BlogRepo.Create(db, obj)
}
