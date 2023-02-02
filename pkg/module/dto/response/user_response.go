package response

type (
	FindAllActiveUserResponse struct {
		Result *ActiveUserResults `json:"result"`

		CodeErr string  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	FindByIDActiveUserResponse struct {
		Result *ActiveUserResult `json:"result"`

		CodeErr string  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	CreateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr string  `json:"code"`
		MsgErr  string `json:"msg"`
	}

	UpdateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr string  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveUserResult struct {
		User *ActiveUserEntity `json:"user"`
	}
	ActiveUserResults struct {
		Users []*ActiveUserEntity `json:"users"`
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
		UserRoll int
		Blogs    []ActiveBlogEntity
		Option   Options
	}
)
