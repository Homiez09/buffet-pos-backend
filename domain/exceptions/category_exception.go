package exceptions

import "errors"

var (
	ErrDuplicatedCategoryName = errors.New("category name already exist")
	ErrCategoryNotFound       = errors.New("category not found")
)
