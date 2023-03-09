package params

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var ValidatorKey = "ValidatorKey"

// FillParams Fill parameters
func FillParams(o interface{}, n interface{}, fields []string) interface{} {
	t := reflect.ValueOf(o).Elem()
	f := reflect.ValueOf(n).Elem()
	for _, field := range fields {
		value := f.FieldByName(field)
		name := t.FieldByName(field)
		if !name.IsValid() {
			continue
		}
		name.Set(value)
	}
	return t
}

// ParamsValidator Verify the parameters
func ParamsValidator(context *gin.Context, params interface{}) error {
	if err := context.ShouldBind(params); err != nil {
		return err
	}
	valid, err := GetValidator(context)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		var e []string
		errs := err.(validator.ValidationErrors)
		for _, v := range errs {
			e = append(e, v.Error())
		}
		return errors.New(strings.Join(e, "\n"))
	}
	return nil
}

// GetValidator Get Params Validator
func GetValidator(context *gin.Context) (*validator.Validate, error) {
	val, ok := context.Get(ValidatorKey)
	if !ok {
		return nil, errors.New("ValidatorDisabled ")
	}
	validate, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("SetValidatorFailed ")
	}
	return validate, nil
}
