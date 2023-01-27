package response

type (
	FindAllActiveUserResponse struct {
		Result *ActiveUserResults `json:"result"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	FindByIDActiveUserResponse struct {
		Result *ActiveUserResult `json:"result"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	CreateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}

	UpdateActiveUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveUserResult struct {
		Student *ActiveUserEntity `json:"student"`
	}
	ActiveUserResults struct {
		Students []*ActiveUserEntity `json:"students"`
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

type ()
