package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	engine *validator.Validate
}

func New() *Validator {
	v := validator.New()
	return &Validator{engine: v}
}

func (v *Validator) Struct(s any) error {
	return v.engine.Struct(s)
}

func (v *Validator) BindAndValidate(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		return err
	}
	return v.Struct(obj)
}

var Validate = New()
