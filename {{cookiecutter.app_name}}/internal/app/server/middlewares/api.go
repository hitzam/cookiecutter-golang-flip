package middlewares

import (
	"mime/multipart"
	"net/http"

	goCoreLog "gitlab.com/flip-id/go-core/helpers/log"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/flip-id/go-core/errors"
	"gitlab.com/flip-id/go-core/helpers/str"
)

// API middleware for API
func API() gin.HandlerFunc {
	return func(c *gin.Context) {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetExtra("method", c.Request.Method)
			scope.SetExtra("url", c.Request.URL.String())
			scope.SetExtra("user_agent", c.Request.UserAgent())
			scope.SetExtra("content_type", c.ContentType())
			scope.SetExtra("query_params", c.Request.URL.Query())
			scope.SetTag("request_id", str.UUID())
			scope.SetExtra("json_response", nil)
			scope.SetUser(sentry.User{ID: ""})
			scope.SetExtra("body_params", nil)
		})

		bodyParams := make(map[string]interface{})
		getBodyParams(c, bodyParams)
		if len(bodyParams) > 0 {
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetExtra("body_params", bodyParams)
			})
		}

		//set request_id - get request-id from header request or create one
		reqID := c.Request.Header.Get(goCoreLog.REQUEST_ID_KEY)
		if reqID == "" {
			reqID = requestid.Get(c)
		}
		c.Set(goCoreLog.REQUEST_ID_KEY, requestid.Get(c))

		c.Next()
	}
}

func getBodyParams(c *gin.Context, bodyParams map[string]interface{}) {
	if c.Request.Method == http.MethodGet || c.Request.Body == http.NoBody {
		return
	}

	b := binding.Default(c.Request.Method, c.ContentType())
	var i interface{} = b
	var err error
	bBody, ok := i.(binding.BindingBody)
	if ok {
		if bBody.Name() == "json" {
			err = c.ShouldBindBodyWith(&bodyParams, bBody) // application/json
		}
	} else if b == binding.FormMultipart {
		err = c.Request.ParseMultipartForm(0) // multipart/form-data
		assignMultipartForm(bodyParams, c.Request.MultipartForm)
	} else if b == binding.Form {
		err = c.Request.ParseForm() // application/x-www-form-urlencoded
		for key, value := range c.Request.PostForm {
			bodyParams[key] = value
		}
	} else {
		err = c.ShouldBind(&bodyParams)
	}

	if err != nil {
		_ = errors.Newf("Get body params error. Error: %v", err)
	}
}

func assignMultipartForm(bodyParams map[string]interface{}, multipartForm *multipart.Form) {
	for key, value := range multipartForm.Value {
		bodyParams[key] = value
	}
	for key, files := range multipartForm.File {
		var fileDetails []map[string]interface{}
		for _, file := range files {
			fileDetail := make(map[string]interface{})
			fileDetail["filename"] = file.Filename
			fileDetail["header"] = file.Header
			fileDetail["filesize"] = file.Size
			fileDetails = append(fileDetails, fileDetail)
		}
		bodyParams[key] = fileDetails
	}
}
