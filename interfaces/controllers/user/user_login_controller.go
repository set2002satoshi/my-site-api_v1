package user

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/auth"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	LoginResponse struct {
		response.LoginUserResponse
	}
)

func (uc *UserController) Login(ctx c.Context) {
	req := &request.UserLoginRequest{}
	res := &LoginResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	loginModel, err := uc.toLoginModel(req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
	}

	createdToken, err := uc.Interactor.FetchToken(loginModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.LoginUserResult{Token: createdToken}
	ctx.JSON(http.StatusOK, res)

}

func (uc *UserController) toLoginModel(req *request.UserLoginRequest) (*auth.UserLoginModel, error) {
	return &auth.UserLoginModel{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func (c *LoginResponse) SetErr(err error) {
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
