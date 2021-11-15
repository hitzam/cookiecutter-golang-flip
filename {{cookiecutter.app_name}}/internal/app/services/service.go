package service

import (
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/repositories"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repositories.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheck
}
