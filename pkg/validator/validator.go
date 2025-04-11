package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("rfc3339date", Date)
}

func Date(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
