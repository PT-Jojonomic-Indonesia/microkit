package validator

import (
	"context"
	"fmt"
	"reflect"
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

	var v int

	switch field.Kind() {
	case reflect.Int:
		v = int(field.Int())
	case reflect.Float32, reflect.Float64:
		vInt, err := strconv.Atoi(fmt.Sprintf("%.0f", field.Float()))
		if err != nil {
			return false
		}
		v = vInt
	}

	if v < 0 {
		panic("negative number")
	}

	n := 0
	for ; v > 0; v /= 10 {
		n += 1
	}

	return n <= param
}
