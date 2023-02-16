package category

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	DeleteByIdActiveCategoryResponse struct {
		response.DeleteByIdCategoryResponse
	}
)

func (cc *CategoryController) Delete(ctx c.Context) {
	req := &request.CategoryDeleteByIdRequest{}
	res := &DeleteByIdActiveCategoryResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	if err := cc.Interactor.Delete(req.ID); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	ctx.JSON(http.StatusOK, "ok")

}

func (c *DeleteByIdActiveCategoryResponse) SetErr(err error) {
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
