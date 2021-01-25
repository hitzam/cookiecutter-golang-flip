package cmd

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/config"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/appcontext"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/commons"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/repository"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/server"
	"github.com/kitabisa/{{ cookiecutter.app_name }}/internal/app/service"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/gorp.v3"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.app_name }}",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		{% if cookiecutter.is_server == "y" -%}
		start()
		{%- endif %}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

{% if cookiecutter.is_server == "y" -%}
func start() {
	cfg := config.Config()

	app := appcontext.NewAppContext(cfg)
	var err error

	var dbMysql *gorp.DbMap
	if cfg.GetBool("MYSQL_IS_ENABLED") {
		dbMysql, err = app.GetDBInstance(appcontext.DBDialectMysql)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
			return
		}
	}

	var dbPostgre *gorp.DbMap
	if cfg.GetBool("POSTGRE_IS_ENABLED") {
		dbPostgre, err = app.GetDBInstance(appcontext.DBDialectPostgres)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Postgre | %v", err)
			return
		}
	}

	var cache *redis.Pool
	if cfg.GetBool("CACHE_IS_ENABLED") {
		cache = app.GetCachePool()
		cacheConn, err := cache.Dial()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Cache | %v", err)
			return
		}
		defer cacheConn.Close()
	}

	var influx *influx.Client
	if cfg.GetBool("INFLUX_IS_ENABLED") {
		influx, err = app.GetInfluxDBClient()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Influx | %v", err)
			return
		}

		err = influx.Ping()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Influx | %v", err)
			return
		}
	}

	opt := commons.Options{
		Config:    cfg,
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
		Influx:    influx,
	}

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	service := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption)

	repo := repository.Repository{
		Cache: cacheRepo,
	}
	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)

	svc := service.Services{
		HealthCheck: hc,
	}
	return &svc
}
{%- endif %}
