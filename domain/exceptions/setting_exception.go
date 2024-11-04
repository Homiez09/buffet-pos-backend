package exceptions

import "errors"

var (
	ErrSettingNotFound      = errors.New("setting key not found")
	ErrDuplicatedSettingKey = errors.New("setting key already exist")
)
