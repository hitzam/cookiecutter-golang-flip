package cmd

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/flip-id/go-core/helpers/array"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/config"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/appcontext"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/driver"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/repositories"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/server"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.app_name }}",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
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
	cfg := config.GetConfig()

	tracer.Start()
	defer tracer.Stop()

	initSentry(cfg)

	app := appcontext.NewAppContext(cfg)
	var err error

	dbMysql, err := app.GetDBInstance(appcontext.DBDialectMysql)
	if err != nil {
		logrus.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
		return
	}

	var cacheClient *redis.Client
	if app.GetCacheOption().IsEnable {
		cacheClient = driver.NewCache(app.GetCacheOption())
		defer cacheClient.Close()
		_, err := cacheClient.Ping(context.Background()).Result()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to Redis | %v", err)
			return
		}
	}

	opt := commons.Options{
		AppCtx:      app,
		CacheClient: cacheClient,
		Config:      cfg,
		DbMysql:     dbMysql,
	}

	repo := wiringRepository(repositories.Option{
		Options: opt,
	})

	service := wiringService(services.Option{
		Options:    opt,
		Repository: repo,
	})

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func initSentry(conf *config.Configuration) {
	if array.InArray(conf.App.Env, []string{"production", "staging"}) && conf.Sentry.DSN != nil {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         *conf.Sentry.DSN,
			Environment: conf.App.Env,
		}); err != nil {
			panic("Cannot initialize sentry")
		}
	} else {
		log.Println("Sentry is not set on this stage")
	}
}

func wiringRepository(repoOption repositories.Option) *repositories.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption)

	repo := repository.Repository{
		Cache: cacheRepo,
	}
	return &repo
}

func wiringService(serviceOption services.Option) *services.Services {
	// wiring up all services
	healthCheck := services.NewHealthCheck(serviceOption)

	svc := services.Services{
		HealthCheck: healthCheck,
	}
	return &svc
}
{%- endif %}
