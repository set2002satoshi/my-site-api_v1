package user

import (
	"github.com/set2002satoshi/my-site-api/interfaces/database"
	DBUser "github.com/set2002satoshi/my-site-api/interfaces/database/user"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/my-site-api/usecase/user"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(db database.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			DB:       &database.DBRepository{DB: db},
			UserRepo: &DBUser.UserRepository{},
		},
	}
}

func (uc *UserController) convertActiveUserToDTO(obj *models.UserEntity) *response.ActiveUserEntity {
	return &response.ActiveUserEntity{
		UserId:   int(obj.GetUserId()),
		UserName: obj.GetUserName(),
		Password: obj.GetPassword(),
		UserRoll: int(obj.GetRoll()),
		Blogs:    []response.ActiveBlogEntity{},
		Option: response.Options{
			// Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}

// func (uc *UserController) convertActiveUserToDTOs(objs []*models.UserEntity) []response.ActiveUserEntity {
// 	UEs := make([]response.ActiveUserEntity, len(objs))
// 	for i, obj := range objs {

// 		result := response.ActiveUserEntity{
// 			UserId:   int(obj.GetUserId()),
// 			UserName: obj.GetUserName(),
// 			Password: obj.GetPassword(),
// 			UserRoll: int(obj.GetUserRoll()),
// 			Blogs:    []response.ActiveBlogEntity{},
// 			Option: response.Options{
// 				// Revision:  int(obj.GetRevision()),
// 				CreatedAt: obj.GetCreatedAt(),
// 				UpdatedAt: obj.GetUpdatedAt(),
// 			},
// 		}
// 		UEs[i] = result
// 	}
// 	return UEs
// }
