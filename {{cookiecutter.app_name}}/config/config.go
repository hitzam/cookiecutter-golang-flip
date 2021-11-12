package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var defaultConfig *viper.Viper

func init() {
}

type AppConfiguration struct {
	Host string
	Env  string
	Port int
	Code string
}

type DatabaseConfiguration struct {
	Driver               string
	Name                 string
	User                 string
	Password             string
	Host                 string
	Port                 int
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

type FlipServerConfiguration struct {
	BaseUrl         string
}
}
type CacheConfiguration struct {
	IsEnable           bool
	Host               string
	Port               int
	DialConnectTimeout time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	Wait               bool
	MaxConnLifetime    time.Duration
	Password           string
	Db                 int
}

type ErrorDesc struct {
	Id string
	En string
}

var configuration *Configuration
var once sync.Once

type Configuration struct {
	App       AppConfiguration
	Database  DatabaseConfiguration
	JWT       JWTConfiguration
	Sentry    SentryConfiguration
	Wappin    WappinConfiguration
	Freshchat FreshchatConfiguration
	Infobip   InfobipConfiguration
	Wavecell  WavecellConfiguration
	Cache     CacheConfiguration
	ErrorMap  map[string]ErrorDesc
}

func GetConfig() *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return configuration
}
