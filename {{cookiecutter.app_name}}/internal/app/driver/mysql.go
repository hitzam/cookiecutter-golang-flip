package driver

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	IsEnable             bool
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

// NewMysqlDatabase return gorm dbmap object with MySQL options param
func NewMysqlDatabase(option DBMysqlOption) (*gorm.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
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

	gorm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	return gorm, nil
}
