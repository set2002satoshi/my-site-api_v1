package category

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	FindByIdActiveCategoryResponse struct {
		response.FindByIdCategoryResponse
	}
)

func (cc *CategoryController) FindById(ctx c.Context) {
	req := &request.CategoryFindByIdRequest{}
	res := &FindByIdActiveCategoryResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	foundCategory, err := cc.Interactor.FindById(req.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveCategoryResult{Category: cc.convertActiveCategoryToDTO(foundCategory)}
	ctx.JSON(http.StatusOK, res)
}

func (bc *FindByIdActiveCategoryResponse) SetErr(err error) {

	h := make([]errors.ErrorInfo, 0)

	for k, v := range errors.ToMap(err) {
		e := errors.ErrorInfo{
			ErrCode: k,
			ErrMsg:  v,
		}
		h = append(h, e)
	}
	bc.Errors = h
}
