package request

type (
	CategoryFindByIdRequest struct {
		ID int
	}
	CategoryDeleteRequest struct {
		ID int
	}
	CategoryCreateRequest struct {
		CategoryName string
	}
)
