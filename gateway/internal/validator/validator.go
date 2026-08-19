package validator

import (
	"time"

	playground "github.com/go-playground/validator/v10"
)

// var Validate = playground.New()
var validate = playground.New()

func init() {
	validate.RegisterValidation("birthdate", validateBirthDate)
}

func validateBirthDate(fl playground.FieldLevel) bool {

	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	min := time.Date(
		1900,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	if date.Before(min) {
		return false
	}

	if date.After(time.Now()) {
		return false
	}

	return true
}

func Validate(v any) error {
	return validate.Struct(v)
}
