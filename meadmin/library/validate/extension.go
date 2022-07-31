package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
)

func datetime(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	valstr := cast.ToString(value)
	err := carbon.ParseByLayout(valstr, carbon.DateTimeLayout).Error
	if err != nil {
		return false
	}
	return true
}

func date(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	valstr := cast.ToString(value)
	err := carbon.ParseByLayout(valstr, carbon.DateLayout).Error
	if err != nil {
		return false
	}
	return true
}

func time(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	valstr := cast.ToString(value)
	err := carbon.ParseByLayout(valstr, carbon.TimeLayout).Error
	if err != nil {
		return false
	}
	return true
}
