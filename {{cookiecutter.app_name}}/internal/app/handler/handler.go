package handler

import (
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/commons"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}
