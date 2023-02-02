package response

type (
	FindAllActiveBlogResponse struct {
		Result *ActiveBlogResults `json:"result"`

		CodeErr error `json:"error"`
		MsgErr  string `json:"msg"`
	}

	FindByIDActiveBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	CreateActiveBlogResponse struct {
		Result *ActiveBlogResult `json:"results"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	UpdateActiveBlogResponse struct {
		Result *ActiveBlogResult `json:"results"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveBlogResult struct {
		Student *ActiveBlogEntity `json:"student"`
	}
	ActiveBlogResults struct {
		Students []*ActiveBlogEntity `json:"students"`
	}

	// HistoryBlogResult struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// HistoryBlogResults struct {
	// 	Students []*HistoryBlogEntity `json:"students"`
	// }

	LoginBlogResult struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

type (
	ActiveBlogEntity struct {
		BlogId  int
		User_id int
		Title   string
		Content string
		Option  Options
	}
)


