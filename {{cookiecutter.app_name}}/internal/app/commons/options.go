package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/appcontext"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/metrics"
	"gopkg.in/gorp.v3"
)

// Options common option for all object that needed
type Options struct {
	AppCtx    *appcontext.AppContext
	Config    config.Provider
	CachePool *redis.Pool
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	Metric    metrics.IMetric
}
