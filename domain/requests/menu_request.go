package requests

import "mime/multipart"

type AddMenuRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description *string        `json:"description" validate:"omitempty"`
	CategoryID  *string        `json:"categoryId" validate:"omitempty"`
	IsAvailable bool           `json:"isAvailable" validate:"required"`
	Price       float64        `json:"price" validate:"required"`
	Image       multipart.File `json:"image" validate:"required"`
}

type MenuFindByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
