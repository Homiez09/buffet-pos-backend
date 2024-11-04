package requests

type AddMenuRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form"description" validate:"omitempty"`
	CategoryID  string `json:"categoryId" form:"categoryId" validate:"omitempty"`
	IsAvailable bool   `json:"isAvailable" form:"isAvailable" validate:"required"`
}

type EditMenuRequest struct {
	ID          string `json:"id" form:"id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"omitempty"`
	Description string `json:"description" form:"description" validate:"omitempty"`
	CategoryID  string `json:"categoryId" form:"categoryId" validate:"omitempty"`
	IsAvailable bool   `json:"isAvailable" form:"isAvailable" validate:"omitempty"`
}

type MenuFindByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
