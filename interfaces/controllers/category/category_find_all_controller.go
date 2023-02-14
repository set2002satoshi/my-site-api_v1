package category

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	FindAllCategoryResponse struct {
		response.FindAllCategoryResponse
	}
)

func (cc *CategoryController) Find(ctx c.Context) {

	res := &FindAllCategoryResponse{}

	founcCategories, err := cc.Interactor.FindAll()
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Results = response.ActiveCategoryResults{Categories: cc.convertActiveCategoryToDTOs(founcCategories)}
	ctx.JSON(http.StatusOK, res)
}

func (cc *FindAllCategoryResponse) SetErr(err error) {

	h := make([]errors.ErrorInfo, 0)

	for k, v := range errors.ToMap(err) {
		e := errors.ErrorInfo{
			ErrCode: k,
			ErrMsg:  v,
		}
		h = append(h, e)
	}
	cc.Errors = h
}
