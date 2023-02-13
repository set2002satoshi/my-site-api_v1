package user

import (
	"time"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	dtoAuth "github.com/set2002satoshi/my-site-api/pkg/module/dto/auth"
	"github.com/set2002satoshi/my-site-api/pkg/module/service/auth"
	"github.com/set2002satoshi/my-site-api/usecase"
	"gorm.io/gorm"
)

type UserInteractor struct {
	DB       usecase.DBRepository
	UserRepo UserRepository
}

func (ui *UserInteractor) FindAll() ([]*models.UserEntity, error) {
	db := ui.DB.Connect()
	return ui.UserRepo.GetAll(db)
}

func (ui *UserInteractor) FindById(id int) (*models.UserEntity, error) {
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

func (ui *UserInteractor) Update(ctx c.Context, obj *models.UserEntity) (*models.UserEntity, error) {
	tx := ui.DB.Begin()
	currentUser, err := ui.UserRepo.GetById(tx, int(obj.GetUserId()))
	if err != nil {
		tx.Rollback()
		return &models.UserEntity{}, err
	}
	if err := currentUser.CountUpRevision(obj.GetRevision()); err != nil {
		tx.Rollback()
		return &models.UserEntity{}, err
	}
	joinObj, err := models.NewUserEntity(
		int(currentUser.GetUserId()),
		obj.GetEmail(),
		obj.GetUserName(),
		obj.GetPassword(),
		string(obj.GetRoll()),
		[]models.BlogEntity{},
		int(currentUser.GetRevision()),
		currentUser.GetCreatedAt(),
		time.Now(),
	)
	if err != nil {
		tx.Rollback()
		return &models.UserEntity{}, err
	}
	updatedObj, err := ui.UserRepo.Update(tx, joinObj)
	if err != nil {
		tx.Rollback()
		return &models.UserEntity{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &models.UserEntity{}, err
	}
	return updatedObj, nil
}

func (ui *UserInteractor) Delete(id int) error {
	db := ui.DB.Connect()
	return ui.UserRepo.Delete(db, id)
}

func (ui *UserInteractor) FetchToken(obj *dtoAuth.UserLoginModel) (string, error) {
	db := ui.DB.Connect()
	findUser, err := ui.UserRepo.FindByEmail(db, obj.Email)
	if err != nil {
		return "", err
	}
	if auth.ComparisonPassAndHash(findUser.GetPassword(), obj.Password) {
		return "", errors.Add(errors.NewCustomError(), errors.SE0002)
	}

	result, err := auth.IssueToken(int(findUser.GetUserId()))
	if err != nil {
		return "", errors.Wrap(errors.NewCustomError(), errors.SE0003, err.Error())
	}
	return result, nil
}

func (ui *UserInteractor) isUniqueEmail(db *gorm.DB, email string) bool {
	count, _ := ui.UserRepo.FetchEmailNumber(db, email)
	return count == 0
}
