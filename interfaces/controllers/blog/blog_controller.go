package blog

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	DBBlog "github.com/set2002satoshi/my-site-api/interfaces/database/blog"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/my-site-api/usecase/blog"
)

type BlogController struct {
	Interactor usecase.BlogInteractor
}

func NewBlogController(db database.DB) *BlogController {
	return &BlogController{
		Interactor: usecase.BlogInteractor{
			DB:       &database.DBRepository{DB: db},
			BlogRepo: &DBBlog.BlogRepository{},
		},
	}
}

func (bc *BlogController) convertActiveBlogToDTO(obj *models.BlogEntity) response.ActiveBlogEntity {
	return response.ActiveBlogEntity{
		BlogId:  int(obj.GetBlogId()),
		UserId:  int(obj.GetUserId()),
		Title:   obj.GetTitle(),
		Content: obj.GetContent(),
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}
