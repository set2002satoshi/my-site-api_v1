package blog

import (
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
	"github.com/set2002satoshi/my-site-api/usecase/user"
)

type BlogInteractor struct {
	DB       usecase.DBRepository
	BlogRepo BlogRepository
	UserRepo user.UserRepository
}

func (bi *BlogInteractor) Register(obj *models.BlogEntity) (*models.BlogEntity, error) {
	db := bi.DB.Connect()
	return bi.BlogRepo.Create(db, obj)
}

func (bi *BlogInteractor) FindById(id int) (*models.UserEntity, error) {
	db := bi.DB.Connect()
	blog, err := bi.BlogRepo.GetById(db, id)
	if err != nil {
		return nil, err
	}
	user, err := bi.UserRepo.GetById(db, int(blog.GetUserId()))
	if err != nil {
		return nil, err
	}
	return bi.blogToUser(user, blog)
}

func (bi *BlogInteractor) blogToUser(user *models.UserEntity, blog models.BlogEntity) (*models.UserEntity, error) {
	BEs := make([]models.BlogEntity, 1)
	BEs[0] = blog

	return models.NewUserEntity(
		int(user.GetUserId()),
		user.GetEmail(),
		user.GetUserName(),
		user.GetPassword(),
		string(user.GetRoll()),
		BEs,
		int(user.GetRevision()),
		user.GetCreatedAt(),
		user.GetUpdatedAt(),
	)
}
