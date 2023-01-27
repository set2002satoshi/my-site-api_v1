package request

type (
	UserFindByIDRequest struct {
		ID int `json:"id"`
	}
	UserDeleteRequest struct {
		ID int `json:"id"`
	}
	UserCreateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Roll     int    `json:"roll"`
	}

	UserUpdateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Roll     int    `json:"roll"`
	}
)