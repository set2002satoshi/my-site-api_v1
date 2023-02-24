package service

import (
	"fmt"
	"time"

	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api/usecase"
	repo "github.com/set2002satoshi/my-site-api/usecase/repository"
	"gorm.io/gorm"
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

	applyBlog, err := models.NewBlogEntity(
		int(obj.GetBlogId()),
		int(obj.GetUserId()),
		userInfo.GetUserName(),
		obj.GetTitle(),
		obj.GetContent(),
		// obj.GetBlogAndCategories(),
		[]models.BlogAndCategoryEntity{},
		[]models.CategoryEntity{},
		int(obj.GetRevision()),
		obj.GetCreatedAt(),
		obj.GetUpdatedAt(),
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	appliedBlog, err := bi.BlogRepo.Create(tx, applyBlog)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = bi.CreateRelatedBlogWithCategories(tx, obj.GetBlogAndCategories(), int(appliedBlog.GetBlogId()))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	appliedCategories, err := bi.RetrieveCategoryEntityFromCategoryIds(tx, obj.GetBlogAndCategories())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	appliedBlogWithCategories, err := models.NewBlogEntity(
		int(appliedBlog.GetBlogId()),
		int(appliedBlog.GetUserId()),
		userInfo.GetUserName(),
		appliedBlog.GetTitle(),
		appliedBlog.GetContent(),
		obj.GetBlogAndCategories(),
		appliedCategories,
		int(appliedBlog.GetRevision()),
		appliedBlog.GetCreatedAt(),
		appliedBlog.GetUpdatedAt(),
	)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &models.BlogEntity{}, err
	}

	fmt.Println(appliedBlogWithCategories)

	return appliedBlogWithCategories, err
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

	blogWithCategoryIds, err := bi.BlogWithCategoryRepo.GetAllByBlogId(db, int(blog.GetBlogId()))
	if err != nil {
		return nil, err
	}

	return bi.blogToUser(db, user, blog, blogWithCategoryIds)
}

func (bi *BlogInteractor) FindAll() ([]*models.BlogEntity, error) {
	db := bi.DB.Connect()
	blogs, err := bi.BlogRepo.GetAll(db)
	if err != nil {
		return nil, err
	}

	blogWithCategoriesEntity := make([]*models.BlogEntity, len(blogs))
	for i, blog := range blogs {
		relatedBlogWithCategory, err := bi.BlogWithCategoryRepo.GetAllByBlogId(db, int(blog.GetBlogId()))
		if err != nil {
			return make([]*models.BlogEntity, 0), err
		}
		categories, err := bi.RetrieveCategoryEntityFromCategoryIds(db, relatedBlogWithCategory)
		if err != nil {
			return make([]*models.BlogEntity, 0), err
		}
		blogWithCategories, err := models.NewBlogEntity(
			int(blog.GetBlogId()),
			int(blog.GetUserId()),
			blog.GetUserName(),
			blog.GetTitle(),
			blog.GetContent(),
			relatedBlogWithCategory,
			categories,
			int(blog.GetRevision()),
			blog.GetCreatedAt(),
			blog.GetUpdatedAt(),
		)
		if err != nil {
			return make([]*models.BlogEntity, 0), err
		}
		blogWithCategoriesEntity[i] = blogWithCategories
	}
	return blogWithCategoriesEntity, nil
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
		[]models.CategoryEntity{},
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

func (bi *BlogInteractor) CreateRelatedBlogWithCategories(tx *gorm.DB, obj []models.BlogAndCategoryEntity, blogId int) (err error) {
	for _, v := range obj {
		ids, err := models.NewBlogAndCategoryEntity(
			types.INITIAL_ID,
			blogId,
			int(v.GetCategoryId()),
		)
		if err != nil {
			return err
		}
		err = bi.BlogWithCategoryRepo.Create(tx, ids)
		if err != nil {
			return err
		}
	}
	return nil
}

// blogWithCategoryのcategoryIDからcategoryEntityを参照しname返す
func (bi *BlogInteractor) RetrieveCategoryEntityFromCategoryIds(tx *gorm.DB, categoryIds []models.BlogAndCategoryEntity) ([]models.CategoryEntity, error) {
	CEs := make([]models.CategoryEntity, len(categoryIds))
	for i, v := range categoryIds {
		ce, err := bi.CategoryRepo.GetById(tx, int(v.GetCategoryId()))
		if err != nil {
			return make([]models.CategoryEntity, 0), err
		}
		CEs[i] = *ce
	}
	return CEs, nil
}

func (bi *BlogInteractor) blogToUser(db *gorm.DB, user *models.UserEntity, blog models.BlogEntity, blogWithCategoryIds []models.BlogAndCategoryEntity) (*models.UserEntity, error) {
	BEs := make([]models.BlogEntity, 1)
	categories, _ := bi.RetrieveCategoryEntityFromCategoryIds(db, blogWithCategoryIds)
	blogWithCategory, err := models.NewBlogEntity(
		int(blog.GetBlogId()),
		int(blog.GetUserId()),
		blog.GetUserName(),
		blog.GetTitle(),
		blog.GetContent(),
		blog.GetBlogAndCategories(),
		categories,
		int(blog.GetRevision()),
		blog.GetCreatedAt(),
		blog.GetUpdatedAt(),
	)
	if err != nil {
		return nil, err
	}
	BEs[0] = *blogWithCategory
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
