package commons

import (
	"github.com/go-redis/redis/v8"
	"github.com/flip-id/{{ cookiecutter.app_name }}/config"
	"github.com/flip-id/{{ cookiecutter.app_name }}/internal/app/appcontext"
	"github.com/flip-id/{{ cookiecutter.app_name }}/internal/app/metrics"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	AppCtx      *appcontext.AppContext
	Config      config.Provider
	CacheClient *redis.Client
	DbMysql     *gorm.DB
	DbPostgre   *gorm.DB
	Metric      metrics.IMetric
}
