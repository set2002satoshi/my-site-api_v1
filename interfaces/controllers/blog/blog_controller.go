package blog

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	"github.com/set2002satoshi/my-site-api/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api/usecase/service"
)

type BlogController struct {
	Interactor service.BlogInteractor
}

func NewBlogController(db config.DB) *BlogController {
	return &BlogController{
		Interactor: service.BlogInteractor{
			DB:                   &config.DBRepository{DB: db},
			BlogRepo:             &database.BlogRepository{},
			UserRepo:             &database.UserRepository{},
			CategoryRepo:         &database.CategoryRepository{},
			BlogWithCategoryRepo: &database.BlogWithCategoryRepository{},
		},
	}
}

func (bc *BlogController) convertActiveBlogToDTO(obj *models.BlogEntity) response.ActiveBlogEntity {

	categoryIds := make([]response.BlogAndCategoryEntity, len(obj.GetBlogAndCategories()))
	categories := make([]response.ActiveCategoryEntity, len(obj.GetBlogAndCategories()))
	for i, c := range obj.GetBlogAndCategories() {
		categoryId := response.BlogAndCategoryEntity{
			Id:         int(c.GetId()),
			BlogId:     int(c.GetBlogId()),
			CategoryId: int(c.GetCategoryId()),
		}
		categoryIds[i] = categoryId
	}
	for i, c := range obj.GetCategories() {
		category := response.ActiveCategoryEntity{
			Id:           int(c.GetCategoryID()),
			CategoryName: c.GetCategoryName(),
		}
		categories[i] = category
	}
	return response.ActiveBlogEntity{
		BlogId:      int(obj.GetBlogId()),
		UserId:      int(obj.GetUserId()),
		UserName:    obj.GetUserName(),
		Title:       obj.GetTitle(),
		Content:     obj.GetContent(),
		Categories:  categories,
		CategoryIds: categoryIds,
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}

func (bc *BlogController) convertActiveBlogToDTOs(obj []*models.BlogEntity) []response.ActiveBlogEntity {
	be := make([]response.ActiveBlogEntity, len(obj))
	for i, bl := range obj {
		categoryIds := make([]response.BlogAndCategoryEntity, len(bl.GetBlogAndCategories()))
		categories := make([]response.ActiveCategoryEntity, len(bl.GetCategories()))
		for i, c := range bl.GetBlogAndCategories() {
			categoryId := response.BlogAndCategoryEntity{
				Id:         int(c.GetId()),
				BlogId:     int(c.GetBlogId()),
				CategoryId: int(c.GetCategoryId()),
			}
			categoryIds[i] = categoryId
		}
		for i, c := range bl.GetCategories() {
			category := response.ActiveCategoryEntity{
				Id:           int(c.GetCategoryID()),
				CategoryName: c.GetCategoryName(),
			}
			categories[i] = category
		}
		blogTmp := response.ActiveBlogEntity{
			BlogId:      int(bl.GetBlogId()),
			UserId:      int(bl.GetUserId()),
			UserName:    bl.GetUserName(),
			Title:       bl.GetTitle(),
			Content:     bl.GetContent(),
			Categories:  categories,
			CategoryIds: categoryIds,
			Option: response.Options{
				Revision:  int(bl.GetRevision()),
				CreatedAt: bl.GetCreatedAt(),
				UpdatedAt: bl.GetUpdatedAt(),
			},
		}
		be[i] = blogTmp
	}
	return be
}

func (bc *BlogController) convertActiveUserWithBlogToDTO(obj *models.UserEntity) response.ActiveUserEntities {
	be := make([]response.ActiveBlogEntity, len(obj.GetBlogs()))
	for i, bl := range obj.GetBlogs() {
		categoryIds := make([]response.BlogAndCategoryEntity, len(bl.GetBlogAndCategories()))
		categories := make([]response.ActiveCategoryEntity, len(bl.GetCategories()))
		for i, c := range bl.GetBlogAndCategories() {
			categoryId := response.BlogAndCategoryEntity{
				Id:         int(c.GetId()),
				BlogId:     int(c.GetBlogId()),
				CategoryId: int(c.GetCategoryId()),
			}
			categoryIds[i] = categoryId
		}
		for i, c := range bl.GetCategories() {
			category := response.ActiveCategoryEntity{
				Id:           int(c.GetCategoryID()),
				CategoryName: c.GetCategoryName(),
			}
			categories[i] = category
		}
		blogTmp := response.ActiveBlogEntity{
			BlogId:      int(bl.GetBlogId()),
			UserId:      int(bl.GetUserId()),
			UserName:    bl.GetUserName(),
			Title:       bl.GetTitle(),
			Content:     bl.GetContent(),
			Categories:  categories,
			CategoryIds: categoryIds,
			Option: response.Options{
				Revision:  int(bl.GetRevision()),
				CreatedAt: bl.GetCreatedAt(),
				UpdatedAt: bl.GetUpdatedAt(),
			},
		}
		be[i] = blogTmp
	}
	return response.ActiveUserEntities{
		UserId:   int(obj.GetUserId()),
		UserName: obj.GetUserName(),
		UserRoll: string(obj.GetRoll()),
		Blogs:    be,
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}
