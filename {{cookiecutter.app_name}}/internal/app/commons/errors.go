package commons

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/flip-id/{{ cookiecutter.app_name }}/config"
	coreHttp "gitlab.com/flip-id/go-core/http"
	"gitlab.com/flip-id/go-core/structs"
)

var cfg = config.Config()
var serviceCode = cfg.App.Code

type Response struct {
	ResponseCode string
	ResponseDesc string
}

// InjectErrors injecting all error response to the handler context
func InjectErrors(handlerCtx *coreHttp.HttpHandlerContext) {
	handlerCtx.AddError(ErrDBConn, ErrDBConnResp)
	handlerCtx.AddError(ErrCacheConn, ErrCacheConnResp)
	// etc...
}

// getErrorResponse will return error response code & description object according to error code
func getErrorResponse(errorCode string) structs.Response {
	res := structs.Response{
		ResponseCode: fmt.Sprintf("%s%s", serviceCode, errorCode),
		ResponseDesc: structs.ResponseDesc{
			ID: cfg.ErrorMap[errorCode].Id,
			EN: cfg.ErrorMap[errorCode].En,
		},
	}
	fmt.Println(cfg.ErrorMap[errorCode].Id)
	fmt.Printf("%+v", res)
	return res
}

// ErrDBConn error type for Error DB Connection
var (
	ErrDBConn = errors.New("ErrDBConn")

	ErrDBConnResp *structs.ErrorResponse = &structs.ErrorResponse{
		Response:   getErrorResponse("1001"),
		HttpStatus: http.StatusInternalServerError,
	}
)

// ErrCacheConn error type for Error Cache Connection
var (
	ErrCacheConn = errors.New("ErrCacheConn")

	ErrCacheConnResp *structs.ErrorResponse = &structs.ErrorResponse{
		Response:   getErrorResponse("1002"),
		HttpStatus: http.StatusInternalServerError,
	}
)
