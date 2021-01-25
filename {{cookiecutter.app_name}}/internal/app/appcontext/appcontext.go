package appcontext

import (
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/driver"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
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

// GetDBInstance getting gorp instance, param: dbType can be "mysql" or "postgre"
func (a *AppContext) GetDBInstance(dbType string) (*gorp.DbMap, error) {
	var gorp *gorp.DbMap
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.getMysqlOption()
		gorp, err = driver.NewMysqlDatabase(dbOption)
	case DBDialectPostgres:
		dbOption := a.getPostgreOption()
		gorp, err = driver.NewPostgreDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gorp, err
}

func (a *AppContext) getMysqlOption() driver.DBMysqlOption {
	return driver.DBMysqlOption{
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

func (a *AppContext) getPostgreOption() driver.DBPostgreOption {
	return driver.DBPostgreOption{
		Host:        a.config.GetString("POSTGRE_HOST"),
		Port:        a.config.GetInt("POSTGRE_PORT"),
		Username:    a.config.GetString("POSTGRE_USERNAME"),
		Password:    a.config.GetString("POSTGRE_PASSWORD"),
		DBName:      a.config.GetString("POSTGRE_DB_NAME"),
		MaxPoolSize: a.config.GetInt("POSTGRE_POOL_SIZE"),
	}
}

// GetCachePool get cache pool connection
func (a *AppContext) GetCachePool() *redis.Pool {
	return driver.NewCache(a.getCacheOption())
}

func (a *AppContext) getCacheOption() driver.CacheOption {
	return driver.CacheOption{
		Host:               a.config.GetString("CACHE_HOST"),
		Port:               a.config.GetInt("CACHE_PORT"),
		Namespace:          a.config.GetString("CACHE_NAMESPACE"),
		Password:           a.config.GetString("CACHE_PASSWORD"),
		DialConnectTimeout: a.config.GetDuration("CACHE_DIAL_CONNECT_TIMEOUT"),
		ReadTimeout:        a.config.GetDuration("CACHE_READ_TIMEOUT"),
		WriteTimeout:       a.config.GetDuration("CACHE_WRITE_TIMEOUT"),
		IdleTimeout:        a.config.GetDuration("CACHE_IDLE_TIMEOUT"),
		MaxConnLifetime:    a.config.GetDuration("CACHE_CONNECTION_MAX_LIFETIME"),
		MaxIdle:            a.config.GetInt("CACHE_MAX_IDLE_CONNECTION"),
		MaxActive:          a.config.GetInt("CACHE_MAX_ACTIVE_CONNECTION"),
		Wait:               a.config.GetBool("CACHE_IS_WAIT"),
	}
}

// GetInfluxDBClient get Influx DB client
func (a *AppContext) GetInfluxDBClient() (c *influx.Client, err error) {
	influxConfig := influx.ClientConfig{
		Addr:               a.config.GetString("INFLUX_HOST"),
		Username:           a.config.GetString("INFLUX_USERNAME"),
		Password:           a.config.GetString("INFLUX_PASSWORD"),
		Database:           a.config.GetString("INFLUX_DB_NAME"),
		RetentionPolicy:    a.config.GetString("INFLUX_RETENTION_POLICY"),
		Timeout:            a.config.GetDuration("INFLUX_TIMEOUT"),
		InsecureSkipVerify: a.config.GetBool("INFLUX_INSECURE_SKIP_VERIFY"),
	}

	return influx.NewClient(influxConfig)
}
