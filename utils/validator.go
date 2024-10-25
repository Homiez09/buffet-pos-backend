package utils

import (
	"errors"
	"fmt"
	"strings"

	valid "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = valid.New()

type ValidateError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ValidateStruct[T any](payload T) *ValidateError {
	err := validate.Struct(payload)
	errMsg := ""
	if err != nil {
		for _, err := range err.(valid.ValidationErrors) {
			tmp := strings.Split(err.StructNamespace(), ".")
			msg := fmt.Sprintf("%s is %s", tmp[len(tmp)-1], err.Tag())
			msg = strings.ToLower(string(msg[0])) + msg[1:]
			errMsg = errMsg + msg + ", "
		}

		return &ValidateError{
			Error:   "Invalid request",
			Message: errMsg[:len(errMsg)-2],
		}
	}

	return nil
}

func ValidateUUID(input string) (*string, error) {
	_, err := uuid.Parse(input)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	return &input, nil
}
