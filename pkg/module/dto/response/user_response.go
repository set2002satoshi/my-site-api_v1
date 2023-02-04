package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllActiveUserResponse struct {
		Results ActiveUserResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIdActiveUserResponse struct {
		Result ActiveUserResult `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActiveUserResponse struct {
		Result ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveUserResponse struct {
		Result ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
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
		UserId   int `json:"user_id"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		UserRoll string `json:"user_roll"`
		Blogs    []ActiveBlogEntity
		Option   Options `json:"option"`
	}
)