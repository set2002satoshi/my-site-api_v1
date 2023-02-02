package user

import (
	"net/http"
	"time"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	CreateActiveUserResponse struct {
		response.CreateActiveUserResponse
	}
)

func (uc *UserController) Create(ctx c.Context) {
	req := &request.UserCreateRequest{}
	res := &CreateActiveUserResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}
	// skip Validation

	reqModel, err := uc.cToModel(req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}
	ok, err := uc.Interactor.Register(reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = &response.ActiveUserResult{User: uc.convertActiveUserToDTO(ok)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *UserController) cToModel(req *request.UserCreateRequest) (*models.UserEntity, error) {
	return models.NewUserEntity(
		types.INITIAL_ID,
		req.Email,
		req.Name,
		req.Password,
		req.Roll,
		time.Time{},
		time.Time{},
	)
}

func (c *CreateActiveUserResponse) SetErr(err error) {

	h := make([]errors.ErrorInfo, 0)

	for k, v := range errors.ToMap(err) {
		e := errors.ErrorInfo{
			ErrCode: k,
			ErrMsg:  v,
		}
		h = append(h, e)
	}
	c.Errors = h
}
