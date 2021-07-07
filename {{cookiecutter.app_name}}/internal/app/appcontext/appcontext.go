package appcontext

import (
	"errors"

	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/driver"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/metrics"
	"gopkg.in/gorp.v3"
)

const (
	// DBDialectMysql rdbms dialect name for MySQL
	DBDialectMysql = "mysql"

	// DBDialectPostgres rdbms dialect name for PostgreSQL
	DBDialectPostgres = "postgres"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetAppOption returns application options
func (a *AppContext) GetAppOption() AppOption {
	return AppOption{
		Host:   a.config.GetString("APP_HOST"),
		Port:   a.config.GetInt("APP_PORT"),
		Name:   a.config.GetString("APP_NAME"),
		Secret: a.config.GetString("APP_SECRET"),
	}
}

// GetDBInstance getting gorp instance, param: dbType can be "mysql" or "postgre"
func (a *AppContext) GetDBInstance(dbType string) (*gorp.DbMap, error) {
	var gorp *gorp.DbMap
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.GetMysqlOption()
		gorp, err = driver.NewMysqlDatabase(dbOption)
	case DBDialectPostgres:
		dbOption := a.GetPostgreOption()
		gorp, err = driver.NewPostgreDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gorp, err
}

// GetMysqlOption returns mysql options
func (a *AppContext) GetMysqlOption() driver.DBMysqlOption {
	return driver.DBMysqlOption{
		IsEnable:             a.config.GetBool("MYSQL_IS_ENABLED"),
		Host:                 a.config.GetString("MYSQL_HOST"),
		Port:                 a.config.GetInt("MYSQL_PORT"),
		Username:             a.config.GetString("MYSQL_USERNAME"),
		Password:             a.config.GetString("MYSQL_PASSWORD"),
		DBName:               a.config.GetString("MYSQL_DB_NAME"),
		AdditionalParameters: a.config.GetString("MYSQL_ADDITIONAL_PARAMS"),
		MaxOpenConns:         a.config.GetInt("MYSQL_MAX_OPEN_CONNECTION"),
		MaxIdleConns:         a.config.GetInt("MYSQL_MAX_IDLE_CONNECTION"),
		ConnMaxLifetime:      a.config.GetDuration("MYSQL_CONNECTION_MAX_LIFETIME"),
	}
}

// GetPostgreOption returns postgresql option
func (a *AppContext) GetPostgreOption() driver.DBPostgreOption {
	return driver.DBPostgreOption{
		IsEnable:    	 a.config.GetBool("POSTGRE_IS_ENABLED"),
		Host:        	 a.config.GetString("POSTGRE_HOST"),
		Port:        	 a.config.GetInt("POSTGRE_PORT"),
		Username:    	 a.config.GetString("POSTGRE_USERNAME"),
		Password:    	 a.config.GetString("POSTGRE_PASSWORD"),
		DBName:      	 a.config.GetString("POSTGRE_DB_NAME"),
		MaxOpenConn: 	 a.config.GetInt("POSTGRE_MAX_OPEN_CONN"),
		MaxIdleConn: 	 a.config.GetInt("POSTGRE_MAX_IDLE_CONN"),
		ConnMaxLifetime: a.config.GetDuration("POSTGRE_CONN_MAX_LIFETIME"),
	}
}

// GetCacheOption returns redis options
func (a *AppContext) GetCacheOption() driver.CacheOption {
	return driver.CacheOption{
		IsEnable:           a.config.GetBool("CACHE_IS_ENABLED"),
		Host:               a.config.GetString("CACHE_HOST"),
		Port:               a.config.GetInt("CACHE_PORT"),
		Namespace:          a.config.GetInt("CACHE_NAMESPACE"),
		Password:           a.config.GetString("CACHE_PASSWORD"),
		DialConnectTimeout: a.config.GetDuration("CACHE_DIAL_CONNECT_TIMEOUT"),
		ReadTimeout:        a.config.GetDuration("CACHE_READ_TIMEOUT"),
		WriteTimeout:       a.config.GetDuration("CACHE_WRITE_TIMEOUT"),
		IdleTimeout:        a.config.GetDuration("CACHE_IDLE_TIMEOUT"),
		MaxConnLifetime:    a.config.GetDuration("CACHE_CONNECTION_MAX_LIFETIME"),
		MaxIdle:            a.config.GetInt("CACHE_MAX_IDLE_CONNECTION"),
		MaxActive:          a.config.GetInt("CACHE_MAX_ACTIVE_CONNECTION"),
	}
}

// GetTelegrafOption return telegraf options
func (a *AppContext) GetTelegrafOption() metrics.TelegrafOption {
	return metrics.TelegrafOption{
		IsEnabled: a.config.GetBool("TELEGRAF_ENABLE"),
		Host:      a.config.GetString("TELEGRAF_HOST"),
		Port:      a.config.GetInt("TELEGRAF_PORT"),
	}
}
