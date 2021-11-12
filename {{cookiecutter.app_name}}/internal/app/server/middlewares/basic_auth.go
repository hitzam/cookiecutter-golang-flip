package middlewares

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/flip-id/go-core/errors"
	"gitlab.com/flip-id/go-core/helpers/response"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/commons"
	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/internal/app/services"
)

// BasicAuthentication Middleware for checking if client is authenticated
func BasicAuth(service services.IClientService) gin.HandlerFunc {
	ctx := context.Background()
	return func(c *gin.Context) {
		// r := container.InitRequest(c)
		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Basic") {
			response.RespondError(c, errors.Unauthorized.New(""))
			c.Abort()
			return
		}
		basicAuthStr := parts[1]
		user, pass := getUserAndPass(basicAuthStr)
		clientObj, err := service.GetClientByUsernameAndPassword(ctx,user, pass)
		
		if err != nil {
			response.RespondError(c, errors.Unauthorized.New(""))
			c.Abort()
			return
		}

		c.Set(commons.BASIC_AUTH_USER_KEY, clientObj)

		c.Next()
	}
}


// Get username and password from basic auth
func getUserAndPass(basicAuthStr string) (username string, password string) {
	decodedAuth, _ := base64.StdEncoding.DecodeString(basicAuthStr)
	s := strings.Split(string(decodedAuth), ":")

	if len(s) != 2 {
		return "", ""
	}

	return s[0], s[1]
}
