package server

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	goCoreHttp "gitlab.com/flip-id/go-core/http"
	goCoreMiddleware "gitlab.com/flip-id/go-core/middleware"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/configs"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/controllers"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/server/middlewares"

	"gitlab.com/flip-id/go-core/structs"
)

// Router a gin gonic
func Router(opt handler.HandlerOption) *chi.Mux {
	cfg := config.GetConfig()
	handlerCtx := goCoreHttp.NewContextHandler(structs.Meta{
		APIEnv:  cfg.App.Env,
	})
	commons.InjectErrors(&handlerCtx)

	r := gin.Default()
	r.Use(middlewares.API())
	r.Use(requestid.New())
	r.Use(middlewares.LogFormatter())
	r.Use(middlewares.CORS())
	r.Use(goCoreMiddleware.HandleError())

	healthCheckController := controllers.NewHealthCheckController(opt, &handlerCtx)
	
	r.GET("/ping", healthCheckController.Ping)
	r.GET("/health-check", healthCheckController.HealthCheck)

	return r
}
