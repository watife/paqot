package helpers

import (
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"strings"
)

// ValidateInputs validates the inputs
func ValidateInputs(dataSet interface{}) (bool, map[string]string) {

	validate := validator.New()

	err := validate.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//Validation error occurred
		errors := make(map[string]string)

		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "The " + name + " is required"
			case "email":
				errors[name] = "The " + name + " should be a valid email"
			case "eqfield":
				errors[name] = "The " + name + " should be equal to the " + err.Param()
			default:
				errors[name] = "The " + name + " is invalid"
			}
		}

		return false, errors
	}
	return true, nil
}

// ValidationError returns the error if input validation fails
func ValidationError(w http.ResponseWriter, code int, payload interface{}) {
	RespondJSON(w, code, payload)
}
