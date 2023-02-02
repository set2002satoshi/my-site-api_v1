package user

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/usecase"
	"gorm.io/gorm"
)

type UserInteractor struct {
	DB       usecase.DBRepository
	UserRepo UserRepository
}

func (ui *UserInteractor) Register(obj *models.UserEntity) (*models.UserEntity, error) {
	db := ui.DB.Connect()
	if !ui.isUniqueEmail(db, obj.GetEmail()) {
		return &models.UserEntity{}, errors.Add(errors.NewCustomError(), errors.REPO0003)
	}
	return ui.UserRepo.Create(db, obj)
}

func (si *UserInteractor) isUniqueEmail(db *gorm.DB, email string) bool {
	count, _ := si.UserRepo.FetchEmailNumber(db, email)
	return count == 0
}
