package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
	goCoreHttp "gitlab.com/flip-id/go-core/http"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
)

// HealthCheckController object for health check controller
type HealthCheckController struct {
	ControllerOption
	goCoreHttp.HttpHandlerContext
}

func NewHealthCheckController(opt ControllerOption, handlerCtx *goCoreHttp.HttpHandlerContext) *HealthCheckController {
	return &HealthCheckController{
		ControllerOption:   opt,
		HttpHandlerContext: *handlerCtx,
	}
}

// HealthCheck checking if all work well
func (h HealthCheckController) HealthCheck(c *gin.Context) {
	logger := goCoreLog.GetLogger(c)

	err := h.Services.HealthCheck.HealthCheckDbMysql(c)
	if err != nil {
		logger.Error(err)
		h.WriteError(c, commons.ErrDBConn)
		return
	}

	if h.AppCtx.GetCacheOption().IsEnable {
		err := h.Services.HealthCheck.HealthCheckDbCache(c)
		if err != nil {
			logger.Error(err)
			h.WriteError(c, commons.ErrCacheConn)
		}
		return
	}
	logger.Error(commons.ErrDBConn.Error())

	h.Write(c, gin.H{"status": "ok"})
}

// Ping ping to server
func (h HealthCheckController) Ping(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}