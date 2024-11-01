package requests

type AddCategoryRequest struct {
	CategoryName string `json:"categoryName" validate:"required"`
}
