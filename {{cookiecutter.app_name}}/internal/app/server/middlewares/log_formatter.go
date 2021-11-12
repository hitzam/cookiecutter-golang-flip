package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LogFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		if param.ErrorMessage == "" {
			return fmt.Sprintf("[%s] \"%s %s %s %d\" [%s] - %s\n",
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
			)
		} else {
			return fmt.Sprintf("[%s] \"%s %s %s %d\" [%s] - %s (Error: %s)\n",
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.ErrorMessage,
			)
		}
	})
}
