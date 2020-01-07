package service

import (
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/commons"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/repository"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheck
}