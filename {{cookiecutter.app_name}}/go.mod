module github.com/flip-id/{{ cookiecutter.app_name }}

go 1.13

require (
	github.com/DataDog/datadog-go v4.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-redis/redis/v8 v8.4.11
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/kitabisa/perkakas/v2 v2.24.0
	github.com/rs/zerolog v1.19.0
	github.com/rubenv/sql-migrate v0.0.0-20210614095031-55d5740dbbcc
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0
	github.com/ziutek/mymysql v1.5.4 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.12
)
