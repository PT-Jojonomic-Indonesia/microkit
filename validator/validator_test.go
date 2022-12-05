package validator_test

import (
	"context"
	"log"
	"testing"

	. "github.com/PT-Jojonomic-Indonesia/microkit/validator"
)

func TestIDlanguage(t *testing.T) {
	type Address struct {
		Street   string `validate:"required"`
		CityName string `validate:"required"`
		Planet   string `validate:"required,lte=1"`
		Phone    string `validate:"required"`
		Distance int32  `validate:"required,int_lte=1"`
	}

	testData := Address{
		Planet:   "bumi",
		Distance: 2233,
	}
	Init()

	err := Validate(context.Background(), testData)
	mapError := GetErrors(err)
	log.Printf("Errors : %+v", mapError)
}
