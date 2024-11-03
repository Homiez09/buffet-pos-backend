package requests

type AddMenuRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form"description" validate:"omitempty"`
	CategoryID  string `json:"categoryId" form:"categoryId" validate:"omitempty"`
	IsAvailable bool   `json:"isAvailable" form:"isAvailable" validate:"required"`
}

type MenuFindByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
