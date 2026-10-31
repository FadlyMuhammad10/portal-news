package request

type CategoryCreateRequest struct {
	Title string `json:"title" validate:"required"`
}
