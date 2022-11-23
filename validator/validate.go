package validator

import (
	"context"
	"strconv"

	validator_lib "github.com/go-playground/validator/v10"
)

var validate *validator_lib.Validate

func Init() {
	validate = validator_lib.New()
	validate.RegisterValidation("int_lte", validateNumberOfDigit)
}

func Validate(ctx context.Context, data interface{}) error {
	if validate == nil {
		Init()
	}
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

func validateNumberOfDigit(fl validator_lib.FieldLevel) bool {
	field := fl.Field()
	param, err := strconv.Atoi(fl.Param())
	if err != nil {
		panic(err.Error())
	}

	v := field.Int()
	if v < 0 {
		panic("negative number")
	}

	n := 0
	for ; v > 0; v /= 10 {
		n += 1
	}

	return n <= param
}
