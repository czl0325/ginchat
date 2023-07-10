package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func PhoneValidator(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if ok && len(value) == 11 {
		regex := regexp.MustCompile("^1[3-9]{1}\\d{9}$")
		result := regex.MatchString(value)
		return result
	}
	return false
}
