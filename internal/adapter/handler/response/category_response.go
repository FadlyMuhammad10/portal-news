package response

type SuccessCategoryResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	CreatedName string `json:"created_name"`
}

type SuccessCreateCategory struct {
	ID int64 `json:"id"`
}
