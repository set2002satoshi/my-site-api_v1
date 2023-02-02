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

func (uc *CreateActiveUserResponse) setErr(err error, errMsg string) {
	uc.CodeErr = err.Error()
	uc.MsgErr = errMsg
}

func (uc *UserController) Create(ctx c.Context) {
	req := &request.UserCreateRequest{}
	res := &CreateActiveUserResponse{}

	if err := ctx.BindJSON(req); err != nil {
		res.setErr(err, errors.ERR0001)
		ctx.JSON(http.StatusOK, res)
		return
	}
	// skip Validation

	reqModel, err := uc.cToModel(req)
	if err != nil {
		res.setErr(err, errors.ERR0000)
		ctx.JSON(http.StatusOK, res)
		return
	}
	ok, err := uc.Interactor.Register(reqModel)
	if err != nil {
		res.setErr(err, errors.ERR0001)
		ctx.JSON(http.StatusOK, res)
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
