package commons

import (
	"gitlab.com/go-redis/redis/v8"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/config"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/appcontext"
	flipserver "gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/pkg/clients/flip_server"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	AppCtx      		*appcontext.AppContext
	Config           	*config.Configuration
	CacheClient 		*redis.Client
	DbMysql     		*gorm.DB
	FlipServerClient 	flipserver.IFlipServerClient
}
