package blog

import (
	"net/http"

	c "github.com/set2002satoshi/my-site-api/interfaces/controllers"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/dto/response"
)

type (
	FindAllActiveBlogResponse struct {
		response.FindAllActiveBlogResponse
	}
)

func (bc *BlogController) Find(ctx c.Context) {

	res := &FindAllActiveBlogResponse{}

	foundBlog, err := bc.Interactor.FindAll()
	if err != nil {
		ctx.JSON(http.StatusOK, errors.Response(errors.Wrap(errors.NewCustomError(), errors.ERR0000, err.Error()), res))
		return
	}
	res.Result = response.ActiveBlogResults{Blogs: bc.convertActiveBlogToDTOs(foundBlog)}
	ctx.JSON(http.StatusOK, res)
}

func (bc *FindAllActiveBlogResponse) SetErr(err error) {

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
