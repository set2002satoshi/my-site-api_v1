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


func(ui *UserInteractor) FindAll() ([]*models.UserEntity, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.GetAll(db)
}

func(ui *UserInteractor) FindById(id int) (*models.UserEntity, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.GetById(db, id)
}

func (ui *UserInteractor) Register(obj *models.UserEntity) (*models.UserEntity, error) {
	db := ui.DB.Connect()
	if !ui.isUniqueEmail(db, obj.GetEmail()) {
		return &models.UserEntity{}, errors.Add(errors.NewCustomError(), errors.REPO0003)
	}
	return ui.UserRepo.Create(db, obj)
}

func (ui *UserInteractor) isUniqueEmail(db *gorm.DB, email string) bool {
	count, _ := ui.UserRepo.FetchEmailNumber(db, email)
	return count == 0
}
