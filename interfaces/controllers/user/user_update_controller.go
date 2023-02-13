package user

import (
	"net/http"
	"time"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	UpdateActiveUserResponse struct {
		response.UpdateActiveUserResponse
	}
)

func (uc *UserController) Update(ctx c.Context) {

	req := &request.UserUpdateRequest{}
	res := &UpdateActiveUserResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	reqModel, err := uc.uToModel(req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
		return
	}

	updatedUser, err := uc.Interactor.Update(ctx, reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveUserWithBlogResults{User: uc.convertActiveUserWithBlogToDTO(updatedUser)}
	ctx.JSON(http.StatusOK, res)
}

func (uc *UserController) uToModel(req *request.UserUpdateRequest) (*models.UserEntity, error) {
	return models.NewUserEntity(
		req.ID,
		req.Email,
		req.Name,
		req.Password,
		req.Roll,
		[]models.BlogEntity{},
		req.Revision,
		time.Time{},
		time.Time{},
	)
}

func (c *UpdateActiveUserResponse) SetErr(err error) {
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
