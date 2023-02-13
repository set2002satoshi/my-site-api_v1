package category

import (
	"fmt"
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	CreateActiveCategoryResponse struct {
		response.CreateCategoryResponse
	}
)

func (cc *CategoryController) Create(ctx c.Context) {
	req := &request.CategoryCreateRequest{}
	res := &CreateActiveCategoryResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	fmt.Println(req)
	reqModel, err := cc.cToModel(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
		return
	}
	created, err := cc.Interactor.Register(reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveCategoryResult{Category: cc.convertActiveCategoryToDTO(created)}
	ctx.JSON(http.StatusOK, res)
}

func (cc *CategoryController) cToModel(ctx c.Context, req *request.CategoryCreateRequest) (*models.CategoryEntity, error) {
	return models.NewCategoryEntity(
		types.INITIAL_ID,
		req.CategoryName,
	)
}

func (bc *CreateActiveCategoryResponse) SetErr(err error) {

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
