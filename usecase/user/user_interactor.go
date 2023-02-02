package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
)

type UserInteractor struct {
	DB       usecase.DBRepository
	UserRepo UserRepository
}

func (ui *UserInteractor) Register(obj *models.UserEntity) (*models.UserEntity, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.Create(db, obj)
}
