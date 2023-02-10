package blog

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	DBBlog "github.com/set2002satoshi/my-site-api/interfaces/database/blog"
	DBUser "github.com/set2002satoshi/my-site-api/interfaces/database/user"
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
			UserRepo: &DBUser.UserRepository{},
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

func (bc *BlogController) convertActiveUserWithBlogToDTOs(obj *models.UserEntity) response.ActiveUserEntities {
	be := make([]response.ActiveBlogEntity, len(obj.GetBlogs()))
	for i, bl := range obj.GetBlogs() {
		blogTmp := response.ActiveBlogEntity{
			BlogId:  int(bl.GetBlogId()),
			UserId:  int(bl.GetUserId()),
			Title:   bl.GetTitle(),
			Content: bl.GetContent(),
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
