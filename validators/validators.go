package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCustomTitle(field validator.FieldLevel) bool {
	return strings.Contains(strings.ToLower(field.Field().String()), "title")
}
