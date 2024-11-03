package exceptions

import "errors"

var (
	ErrMenuNotFound       = errors.New("menu not found")
	ErrDuplicatedMenuName = errors.New("menu name already exist")
)
