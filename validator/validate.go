package validator

import (
	"context"

	validator_lib "github.com/go-playground/validator/v10"
)

var validate *validator_lib.Validate

func Init() {
	validate = validator_lib.New()
}

func Validate(ctx context.Context, data interface{}) error {
	return validate.StructCtx(ctx, data)
}

func GetErrors(err error) map[string]interface{} {
	mapErrors := make(map[string]interface{})
	if err != nil {
		errs := err.(validator_lib.ValidationErrors)
		for _, e := range errs {
			mapErrors[e.Field()] = e.Error()
		}
	}
	return mapErrors
}
