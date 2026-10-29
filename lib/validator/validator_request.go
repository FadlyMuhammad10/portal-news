package validatorLib

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	var errMsg []string

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errMsg = append(errMsg, "Invalid email format")
			case "required":
				errMsg = append(errMsg, "Field"+err.Field()+" must be required")
			case "min":
				if err.Field() == "Password" {
					errMsg = append(errMsg, "Password minimum 8 character")
				}
			default:
				errMsg = append(errMsg, "Field "+err.Field()+" not valid")
			}
		}
		return errors.New("Validation failed: " + joinMessage(errMsg))
	}
	return nil
}

func joinMessage(msgs []string) string {
	result := ""
	for i, msg := range msgs {
		if i > 0 {
			result += ", "
		}
		result += msg
	}

	return result
}
