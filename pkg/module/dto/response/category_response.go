package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllCategoryResponse struct {
		Results ActiveCategoryResults
		Errors  []errors.ErrorInfo
	}
	FindByIdCategoryResponse struct {
		Result ActiveCategoryResult
		Errors []errors.ErrorInfo
	}
	CreateCategoryResponse struct {
		Result ActiveCategoryResult
		Errors []errors.ErrorInfo
	}

	DeleteCategoryResponse struct {
		Errors []errors.ErrorInfo
	}
)

type (
	ActiveCategoryResult struct {
		Blog ActiveBlogEntity `json:"blog"`
	}
	ActiveCategoryResults struct {
		Blogs []ActiveBlogEntity `json:"blogs"`
	}
)

type (
	ActiveCategoryEntity struct {
		Id           int    `json:"id"`
		CategoryName string `json:"category_name"`
	}
)
