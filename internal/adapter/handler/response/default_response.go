package response

type ErrorResponseDefault struct {
	Meta
}
type Meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseDefault struct {
	Meta       Meta               `json:"meta"`
	Data       interface{}        `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	TotalData  int `json:"total_data"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
}

type SuccessCreateResponseDefault struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
