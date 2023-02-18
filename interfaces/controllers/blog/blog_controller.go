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

	categories := make([]response.BlogAndCategoryEntity, len(obj.Categories))
	for i, c := range obj.Categories {
		res := response.BlogAndCategoryEntity{
			Id:         int(c.GetId()),
			BlogId:     int(c.GetBlogId()),
			CategoryId: int(c.GetCategoryId()),
		}
		categories[i] = res
	}

	return response.ActiveBlogEntity{
		BlogId:     int(obj.GetBlogId()),
		UserId:     int(obj.GetUserId()),
		UserName:   obj.GetUserName(),
		Title:      obj.GetTitle(),
		Content:    obj.GetContent(),
		Categories: categories,
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
		categories := make([]response.BlogAndCategoryEntity, len(bl.Categories))
		for i, c := range bl.Categories {
			res := response.BlogAndCategoryEntity{
				Id:         int(c.GetId()),
				BlogId:     int(c.GetBlogId()),
				CategoryId: int(c.GetCategoryId()),
			}
			categories[i] = res
		}
		blogTmp := response.ActiveBlogEntity{
			BlogId:     int(bl.GetBlogId()),
			UserId:     int(bl.GetUserId()),
			UserName:   bl.GetUserName(),
			Title:      bl.GetTitle(),
			Content:    bl.GetContent(),
			Categories: categories,
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
		categories := make([]response.BlogAndCategoryEntity, len(bl.Categories))
		for i, c := range bl.Categories {
			res := response.BlogAndCategoryEntity{
				Id:         int(c.GetId()),
				BlogId:     int(c.GetBlogId()),
				CategoryId: int(c.GetCategoryId()),
			}
			categories[i] = res
		}
		blogTmp := response.ActiveBlogEntity{
			BlogId:     int(bl.GetBlogId()),
			UserId:     int(bl.GetUserId()),
			UserName:   bl.GetUserName(),
			Title:      bl.GetTitle(),
			Content:    bl.GetContent(),
			Categories: categories,
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
