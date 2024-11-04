package exceptions

import "errors"

var (
	ErrDuplicatedTableName  = errors.New("table name already exist")
	ErrTableNotFound        = errors.New("table not found")
	ErrTableAlreadyAssigned = errors.New("table already assigned")
)
