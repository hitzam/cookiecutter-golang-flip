module gitlab.com/flip-id/{{ cookiecutter.app_name }}

go 1.13

require (
	github.com/DataDog/datadog-go v4.4.0+incompatible
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/getsentry/sentry-go v0.4.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/requestid v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.4.11
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/lib/pq v1.10.0
	github.com/rubenv/sql-migrate v0.0.0-20210614095031-55d5740dbbcc
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0
	gitlab.com/flip-id/go-core v0.0.7
	gopkg.in/DataDog/dd-trace-go.v1 v1.33.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)
