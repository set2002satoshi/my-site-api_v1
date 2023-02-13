package request

type (
	BlogFindByIdRequest struct {
		ID int `json:"id"`
	}
	BlogDeleteRequest struct {
		ID int `json:"id"`
	}
	BlogCreateRequest struct {
		Title   string `json:"title"`
		Context string `json:"context"`
	}

	BlogUpdateRequest struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Context  string `json:"context"`
		Revision int    `json:"revision"`
	}
)

type (
	BlogLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"pass"`
	}
)
