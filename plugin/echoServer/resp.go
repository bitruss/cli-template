package echoServer

import (
	"net/http"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/labstack/echo/v4"
)

type RespBody struct {
	Status int         `json:"status" `
	Result interface{} `json:"result" `
	Msg    string      `json:"msg" `
}

//status <0
func ErrorResp(c echo.Context, status int, data interface{}, msg string) error {
	if status >= 0 {
		basic.Logger.Panicln("error response status should be < 0")
	}
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Result: data,
		Msg:    msg,
	})
}

//status >0
func SuccessResp(c echo.Context, status int, data interface{}, msg string) error {
	if status <= 0 {
		basic.Logger.Panicln("success response status should be > 0")
	}
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Result: data,
		Msg:    msg,
	})
}
