package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"gopkg.in/gorp.v3"
)

// Options common option for all object that needed
type Options struct {
	Config    config.Provider
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
}
