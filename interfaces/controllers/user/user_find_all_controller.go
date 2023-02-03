package user

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	FindAllActiveUserResponse struct {
		response.FindAllActiveUserResponse
	}
)

func (uc *UserController) FindAll(ctx c.Context) {

	res := &FindAllActiveUserResponse{}

	userAll, err := uc.Interactor.FindAll()
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Results = response.ActiveUserResults{Users: uc.convertActiveUserToDTOs(userAll)}
	ctx.JSON(http.StatusOK, res)

}

func (c *FindAllActiveUserResponse) SetErr(err error) {
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
