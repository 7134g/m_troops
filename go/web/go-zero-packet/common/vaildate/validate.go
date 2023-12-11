package vaildate

import (
	"context"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Binding struct {
}

func NewValidateMiddleware() *Binding {
	return &Binding{}
}

func (b Binding) Validate(ctx context.Context, params interface{}) error {
	valid := validator.New()
	if err := valid.StructCtx(ctx, params); err != nil {
		return err
	}

	return nil
}

func (b Binding) Success(resp, data interface{}) {
	respValue := reflect.ValueOf(resp).Elem()
	dataValue := reflect.ValueOf(data)

	respValue.FieldByName("Data").Set(dataValue)
	respValue.FieldByName("Code").Set(reflect.ValueOf(200))
}

func (b Binding) Error(resp interface{}, code int, err error) {

	respValue := reflect.ValueOf(resp).Elem()
	respValue.FieldByName("Message").Set(reflect.ValueOf(err.Error()))
	respValue.FieldByName("Code").Set(reflect.ValueOf(code))
}
