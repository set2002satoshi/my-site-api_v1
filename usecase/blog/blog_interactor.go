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
	userInfo, err := bi.UserRepo.GetById(db, int(obj.GetUserId()))
	if err != nil {
		return nil, err
	}
	applyBlog, err := models.NewBlogEntity(
		int(obj.GetBlogId()),
		int(obj.GetUserId()),
		userInfo.GetUserName(),
		obj.GetTitle(),
		obj.GetContent(),
		int(obj.GetRevision()),
		obj.GetCreatedAt(),
		obj.GetUpdatedAt(),
	)
	if err != nil {
		return nil, err
	}
	return bi.BlogRepo.Create(db, applyBlog)
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

func (bi *BlogInteractor) FindAll() ([]*models.BlogEntity, error) {
	db := bi.DB.Connect()
	return bi.BlogRepo.GetAll(db)
}

func (bi *BlogInteractor) Delete(id int) error {
	db := bi.DB.Connect()
	return bi.BlogRepo.Delete(db, id)
}

func (bi *BlogInteractor) blogToUser(user *models.UserEntity, blog models.BlogEntity) (*models.UserEntity, error) {
	BEs := make([]models.BlogEntity, 1)
	BEs[0] = blog

	return models.NewUserEntity(
		int(user.GetUserId()),
		user.GetEmail(),
		user.GetUserName(),
		string(user.GetPassword()),
		string(user.GetRoll()),
		BEs,
		int(user.GetRevision()),
		user.GetCreatedAt(),
		user.GetUpdatedAt(),
	)
}
