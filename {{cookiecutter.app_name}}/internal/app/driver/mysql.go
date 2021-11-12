package driver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	Host                 string
	Port                 int
	User                 string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

// NewMysqlDatabase return gorm dbmap object with MySQL options param
func NewMysqlDatabase(option DBMysqlOption) (*gorm.DB, error) {
	// Register augments the provided driver with tracing, enabling it to be loaded by gormtrace.Open.
	sqltrace.Register("mysql", &pq.Driver{})

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.User, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(option.MaxOpenConns)
	db.SetMaxIdleConns(option.MaxIdleConns)
	db.SetConnMaxLifetime(option.ConnMaxLifetime)

	gorm, err := gormtrace.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	return gorm, nil
}
