package service

import (
	"context"

	"github.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	plog "github.com/kitabisa/perkakas/v2/log"
)

// IHealthCheck interface for health check service
type IHealthCheck interface {
	HealthCheckDbMysql(ctx context.Context) (err error)
	HealthCheckDbPostgres(ctx context.Context) (err error)
	HealthCheckDbCache(ctx context.Context) (err error)
}

type healthCheck struct {
	opt Option
}

// NewHealthCheck create health check service instance with option as param
func NewHealthCheck(opt Option) IHealthCheck {
	return &healthCheck{
		opt: opt,
	}
}

func (h *healthCheck) HealthCheckDbMysql(ctx context.Context) (err error) {
	sqlDB, _ := h.opt.DbMysql.DB()
	err = sqlDB.Ping()
	if err != nil {
		plog.Zlogger(ctx).Err(err).Send()
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbPostgres(ctx context.Context) (err error) {
	sqlDB, _ := h.opt.DbPostgre.DB()
	err = sqlDB.Ping()
	if err != nil {
		plog.Zlogger(ctx).Err(err).Send()
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbCache(ctx context.Context) (err error) {
	err = h.opt.CacheClient.Ping(ctx).Err()
	if err != nil {
		plog.Zlogger(ctx).Err(err).Send()
		err = commons.ErrCacheConn
		return
	}

	return nil
}
