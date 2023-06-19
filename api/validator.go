package api

import (
	"github.com/caard0s0/united-atomic-bank/util"

	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
