package driver

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBPostgreOption options for postgre connection
type DBPostgreOption struct {
	IsEnable        bool
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns     int
	ConnMaxLifetime time.Duration
}

// NewPostgreDatabase return gorm dbmap object with postgre options param
func NewPostgreDatabase(option DBPostgreOption) (*gorm.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", option.Host, option.Port, option.Username, option.DBName, option.Password))
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

	gorm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})


	return gorm, nil
}
