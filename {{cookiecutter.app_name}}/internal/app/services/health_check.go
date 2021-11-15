package services

import (
	"context"

	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
)

// IHealthCheck interface for health check service
type IHealthCheck interface {
	HealthCheckDbMysql(ctx context.Context) (err error)
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

func (svc *healthCheck) HealthCheckDbMysql(ctx context.Context) (err error) {
	sqlDB, _ := svc.opt.DbMysql.DB()
	err = sqlDB.Ping()
	if err != nil {
		goCoreLog.GetLogger(ctx).FormatLog("DB ping", err, "").Error("Failed")
		err = commons.ErrDBConn
	}
	return
}

func (svc *healthCheck) HealthCheckDbCache(ctx context.Context) (err error) {
	err = svc.opt.CacheClient.Ping(ctx).Err()
	if err != nil {
		goCoreLog.GetLogger(ctx).FormatLog("Cache ping", err, "").Error("Failed")
		err = commons.ErrCacheConn
		return
	}

	return nil
}
