package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type errorValidation struct {
	Errors map[string]string `json:"errors"`
}

func ValidateRequest(verr validator.ValidationErrors) errorValidation {
	errs := make(map[string]string)
	for _, f := range verr {
		err := f.ActualTag()
		err = translate(strings.ToLower(f.Field()), err, f.Param())
		errs[strings.ToLower(f.Field())] = err
	}

	return errorValidation{errs}
}

func translate(field, err, param string) string {
	if err == "required" {
		return fmt.Sprint(field + " is required")
	}

	if err == "min" {
		if field == "age" {
			return fmt.Sprintf("%s cannot be lower than %s", field, param)
		}

		return fmt.Sprintf("%s has to have a minimum length of %s characters", field, param)
	}

	return fmt.Sprint(field + " is invalid")
}
