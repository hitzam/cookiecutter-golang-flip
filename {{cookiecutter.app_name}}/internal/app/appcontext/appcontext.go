package appcontext

import (
	"errors"

	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/config"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/driver"
	flipserver "gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/pkg/clients/flip_server"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/metrics"
	"gorm.io/gorm"
)

const (
	// DBDialectMysql rdbms dialect name for MySQL
	DBDialectMysql = "mysql"

	// DBDialectPostgres rdbms dialect name for PostgreSQL
	DBDialectPostgres = "postgres"
)

// AppContext the app context struct
type AppContext struct {
	config *config.Configuration
}

// NewAppContext initiate appcontext object
func NewAppContext(config *config.Configuration) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetAppOption returns application options
func (a *AppContext) GetAppOption() AppOption {
	return AppOption{
		Host: a.config.App.Host,
		Port: a.config.App.Port,
		Env:  a.config.App.Env,
	}
}

// GetDBInstance getting gorm instance, param: dbType can be "mysql" or "postgre"
func (a *AppContext) GetDBInstance(dbType string) (*gorm.DB, error) {
	var gorm *gorm.DB
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.GetMysqlOption()
		gorm, err = driver.NewMysqlDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gorm, err
}

// GetMysqlOption returns mysql options
func (a *AppContext) GetMysqlOption() driver.DBMysqlOption {
	return driver.DBMysqlOption{
		Host:                 a.config.Database.Host,
		Port:                 a.config.Database.Port,
		User:                 a.config.Database.User,
		Password:             a.config.Database.Password,
		DBName:               a.config.Database.Name,
		AdditionalParameters: a.config.Database.AdditionalParameters,
		MaxOpenConns:         a.config.Database.MaxOpenConns,
		MaxIdleConns:         a.config.Database.MaxIdleConns,
		ConnMaxLifetime:      a.config.Database.ConnMaxLifetime,
	}
}

// GetCacheOption returns redis options
func (a *AppContext) GetCacheOption() driver.CacheOption {
	return driver.CacheOption{
		IsEnable:           a.config.Cache.IsEnable,
		Host:               a.config.Cache.Host,
		Port:               a.config.Cache.Port,
		DB:                 a.config.Cache.Db,
		Password:           a.config.Cache.Password,
		DialConnectTimeout: a.config.Cache.DialConnectTimeout,
		ReadTimeout:        a.config.Cache.ReadTimeout,
		WriteTimeout:       a.config.Cache.WriteTimeout,
		IdleTimeout:        a.config.Cache.IdleTimeout,
		MaxConnLifetime:    a.config.Cache.MaxConnLifetime,
		MaxIdle:            a.config.Cache.MaxIdle,
		MaxActive:          a.config.Cache.MaxActive,
	}
}

// GetFlipServerClient return flip_server interface
func (a *AppContext) GetFlipServerClient() flipserver.IFlipServerClient {
	OPT := flipserver.Option{
		BaseUrl:   a.config.FlipServer.BaseUrl,
	}

	return flipserver.NewFlipServerClient(OPT)
}
