package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	valid "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
		// Check if the error is of type ValidationErrors
		if validationErrors, ok := err.(valid.ValidationErrors); ok {
			for _, err := range validationErrors {
				tmp := strings.Split(err.StructNamespace(), ".")
				msg := fmt.Sprintf("%s is %s", tmp[len(tmp)-1], err.Tag())
				msg = strings.ToLower(string(msg[0])) + msg[1:]
				errMsg += msg + ", "
			}
			return &ValidateError{
				Error:   "Invalid request",
				Message: errMsg[:len(errMsg)-2],
			}
		} else if _, ok := err.(*valid.InvalidValidationError); ok {
			// Handle invalid validation errors separately if needed
			return &ValidateError{
				Error:   "Validation error",
				Message: "Invalid input structure",
			}
		}
		// Unexpected error type
		return &ValidateError{
			Error:   "Unexpected error",
			Message: err.Error(),
		}
	}

	return nil
}

func PopulateStructFromForm[T any](c *fiber.Ctx, dst *T) error {
	val := reflect.ValueOf(dst).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		formKey := fieldType.Tag.Get("form")
		if formKey == "" {
			formKey = fieldType.Name
		}

		formValue := c.FormValue(formKey)
		if formValue == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(formValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intValue, err := strconv.Atoi(formValue); err == nil {
				field.SetInt(int64(intValue))
			}
		case reflect.Float32, reflect.Float64:
			if floatValue, err := strconv.ParseFloat(formValue, 64); err == nil {
				field.SetFloat(floatValue)
			}
		case reflect.Bool:
			if boolValue, err := strconv.ParseBool(formValue); err == nil {
				field.SetBool(boolValue)
			}
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

func ValidatePrice(price string) error {
	if _, err := strconv.ParseFloat(price, 64); err != nil {
		return errors.New("invalid price format")
	}
	return nil
}
