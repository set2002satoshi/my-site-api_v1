package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllActiveUserResponse struct {
		Results ActiveUserResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIDActiveUserResponse struct {
		Result *ActiveUserResult `json:"result"`

		Errors errors.Errors
	}

	CreateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr string  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveUserResult struct {
		User ActiveUserEntity `json:"user"`
	}
	ActiveUserResults struct {
		Users []ActiveUserEntity `json:"users"`
	}

	// HistoreUserResult struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// HistoreUserResults struct {
	// 	Students []*HistoryUserEntity `json:"students"`
	// }

	LoginUserResult struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

type (
	ActiveUserEntity struct {
		UserId   int
		UserName string
		Password string
		UserRoll string
		Blogs    []ActiveBlogEntity
		Option   Options
	}
)
		