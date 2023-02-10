package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllActiveBlogResponse struct {
		Result ActiveBlogResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIDActiveBlogResponse struct {
		Result ActiveUserWithBlogResult `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActiveBlogResponse struct {
		Result ActiveBlogResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveBlogResponse struct {
		Result ActiveBlogResult `json:"results"`

		Errors []errors.ErrorInfo
	}
)

type (
	ActiveBlogResult struct {
		Blog ActiveBlogEntity `json:"blog"`
	}
	ActiveUserWithBlogResult struct {
		Blog ActiveUserEntities `json:"user_with_blog"`
	}
	ActiveBlogResults struct {
		Blogs []ActiveBlogEntity `json:"blogs"`
	}

	// HistoryBlogResult struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// HistoryBlogResults struct {
	// 	Students []*HistoryBlogEntity `json:"students"`
	// }

)

type (
	ActiveBlogEntity struct {
		BlogId  int
		UserId  int
		Title   string
		Content string
		Option  Options
	}
)
