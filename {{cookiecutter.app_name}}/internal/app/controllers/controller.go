package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/flip-id/go-core/helpers/method"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/services"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}

func (ctl *ControllerOption) Validate(param interface{}) error {
	err := method.ValidateStruct(param)
	if err != nil {
		return err
	}
	return nil
}

func (ctl *ControllerOption) SetPostParams(c *gin.Context, params interface{}) error {
	if params == nil {
		return nil
	}
	b := binding.Default(c.Request.Method, c.ContentType())
	var i interface{} = b
	var err error
	bBody, ok := i.(binding.BindingBody)
	if ok {
		// Use ShouldBindBodyWith so we can reuse request body after we read it (so we can have multiple binding)
		err = c.ShouldBindBodyWith(params, bBody)
	} else {
		err = c.ShouldBind(params)
	}

	if err != nil {
		_ = c.Error(err)
		return err
	}

	// r.PostParams = params
	return nil
}