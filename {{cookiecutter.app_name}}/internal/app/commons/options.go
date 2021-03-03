package commons

import (
	"github.com/go-redis/redis/v8"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/appcontext"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/metrics"
	"gopkg.in/gorp.v3"
)

// Options common option for all object that needed
type Options struct {
	AppCtx      *appcontext.AppContext
	Config      config.Provider
	CacheClient *redis.Client
	DbMysql     *gorp.DbMap
	DbPostgre   *gorp.DbMap
	Metric      metrics.IMetric
}
