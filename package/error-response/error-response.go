package errorresponse

import "github.com/go-playground/validator/v10"

type ErrorStruct struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func BindingValidator(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is Required!"
	case "email":
		return "Invalid " + fe.Field() + " Format!"
	case "min":
		return fe.Field() + " Must Be At Least " + fe.Param() + " Characters!"
	default:
		return "Invalid " + fe.Field()
	}
}
