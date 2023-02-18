package service

import (
	"time"

	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/usecase"
	repo "github.com/set2002satoshi/my-site-api/usecase/repository"
)

type BlogInteractor struct {
	DB                   usecase.DBRepository
	UserRepo             repo.UserRepository
	BlogRepo             repo.BlogRepository
	CategoryRepo         repo.CategoryRepository
	BlogWithCategoryRepo repo.BlogWithCategoryRepository
}

func (bi *BlogInteractor) Register(obj *models.BlogEntity) (*models.BlogEntity, error) {
	tx := bi.DB.Begin()

	userInfo, err := bi.UserRepo.GetById(tx, int(obj.GetUserId()))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	createdCategories, err := bi.BlogWithCategoryRepo.BatchCreate(tx, obj.GetBlogAndCategories())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	applyBlog, err := models.NewBlogEntity(
		int(obj.GetBlogId()),
		int(obj.GetUserId()),
		userInfo.GetUserName(),
		obj.GetTitle(),
		obj.GetContent(),
		createdCategories,
		int(obj.GetRevision()),
		obj.GetCreatedAt(),
		obj.GetUpdatedAt(),
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = bi.BlogRepo.Create(tx, applyBlog)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	return applyBlog, err
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

func (bi *BlogInteractor) Update(obj *models.BlogEntity) (*models.BlogEntity, error) {
	tx := bi.DB.Begin()
	currentBlog, err := bi.BlogRepo.GetById(tx, int(obj.GetBlogId()))
	if err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	if err := currentBlog.CountUpRevision(obj.GetRevision()); err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	joinObj, err := models.NewBlogEntity(
		int(currentBlog.GetBlogId()),
		int(currentBlog.GetUserId()),
		currentBlog.GetUserName(),
		obj.GetTitle(),
		obj.GetContent(),
		obj.GetBlogAndCategories(),
		int(currentBlog.GetRevision()),
		currentBlog.GetCreatedAt(),
		time.Now(),
	)
	if err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	updatedObj, err := bi.BlogRepo.Update(tx, joinObj)
	if err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}
	return updatedObj, nil
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
