package request

type (
	CategoryFindByIdRequest struct {
		ID int `json:"id"`
	}
	CategoryDeleteRequest struct {
		ID int `json:"id"`
	}
	CategoryCreateRequest struct {
		CategoryName string `json:"category_name"`
	}
)
