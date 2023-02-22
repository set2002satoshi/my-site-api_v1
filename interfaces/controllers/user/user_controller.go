package user

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	"github.com/set2002satoshi/my-site-api/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api/usecase/service"
)

type UserController struct {
	Interactor service.UserInteractor
}

func NewUserController(db config.DB) *UserController {
	return &UserController{
		Interactor: service.UserInteractor{
			DB:       &config.DBRepository{DB: db},
			UserRepo: &database.UserRepository{},
		},
	}
}

func (uc *UserController) convertActiveUserToDTO(obj *models.UserEntity) response.ActiveUserEntity {
	return response.ActiveUserEntity{
		UserId:   int(obj.GetUserId()),
		UserName: obj.GetUserName(),
		Email:    obj.GetEmail(),
		Password: obj.GetPassword(),
		UserRoll: string(obj.GetRoll()),
		Blog:     []response.ActiveBlogEntity{},
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}

func (uc *UserController) convertActiveUserToDTOs(objs []*models.UserEntity) []response.ActiveUserEntity {
	UEs := make([]response.ActiveUserEntity, len(objs))
	for i, obj := range objs {

		result := response.ActiveUserEntity{
			UserId:   int(obj.GetUserId()),
			UserName: obj.GetUserName(),
			Email:    obj.GetEmail(),
			Password: obj.GetPassword(),
			UserRoll: string(obj.GetRoll()),
			Blog:     []response.ActiveBlogEntity{},
			Option: response.Options{
				Revision:  int(obj.GetRevision()),
				CreatedAt: obj.GetCreatedAt(),
				UpdatedAt: obj.GetUpdatedAt(),
			},
		}
		UEs[i] = result
	}
	return UEs
}

func (bc *UserController) convertActiveUserWithBlogToDTO(obj *models.UserEntity) response.ActiveUserEntities {
	be := make([]response.ActiveBlogEntity, len(obj.GetBlogs()))
	for i, bl := range obj.GetBlogs() {
		blogTmp := response.ActiveBlogEntity{
			BlogId:   int(bl.GetBlogId()),
			UserId:   int(bl.GetUserId()),
			UserName: bl.GetUserName(),
			Title:    bl.GetTitle(),
			Content:  bl.GetContent(),
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
		Email:    obj.GetEmail(),
		Password: obj.GetPassword(),
		UserRoll: string(obj.GetRoll()),
		Blogs:    be,
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}
