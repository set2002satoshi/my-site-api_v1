package blog

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/models"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/request"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	CreateActiveBlogResponse struct {
		response.CreateActiveBlogResponse
	}
)

func (bc *BlogController) Create(ctx c.Context) {
	req := &request.BlogCreateRequest{}
	res := &CreateActiveBlogResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0001, err.Error()), res))
		return
	}

	reqModel, err := bc.cToModel(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0002, err.Error()), res))
		return
	}
	fmt.Println(reqModel.GetBlogAndCategories())
	createdBlog, err := bc.Interactor.Register(reqModel)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveBlogResult{Blog: bc.convertActiveBlogToDTO(createdBlog)}
	ctx.JSON(http.StatusOK, res)
}

func (bc *BlogController) cToModel(ctx c.Context, req *request.BlogCreateRequest) (*models.BlogEntity, error) {

	userSId, isNoErr := ctx.Get("userID")
	if !isNoErr {
		return &models.BlogEntity{}, errors.Add(errors.NewCustomError(), errors.ERR0003)
	}
	userId, _ := strconv.Atoi(userSId.(string))

	categories := make([]models.BlogAndCategoryEntity, len(req.Categories))
	for i, c := range req.Categories {
		category, err := models.NewBlogAndCategoryEntity(
			types.INITIAL_ID,
			types.INITIAL_ID,
			c.ID,
		)
		if err != nil {
			return &models.BlogEntity{}, err
		}
		categories[i] = *category
	}

	return models.NewBlogEntity(
		types.INITIAL_ID,
		userId,
		types.DEFAULT_NAME,
		req.Title,
		req.Context,
		categories,
		[]models.CategoryEntity{},
		types.INITIAL_REVISION,
		time.Time{},
		time.Time{},
	)
}

func (bc *CreateActiveBlogResponse) SetErr(err error) {

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
