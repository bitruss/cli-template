package echoServer

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RespBody struct {
	Status int         `json:"status" `
	Result interface{} `json:"result" `
	Msg    string      `json:"msg" `
}

//status <-1
func ErrorResp(c echo.Context, status int, msg string) error {
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Msg:    msg,
	})
}

//status >=0
func SuccessResp(c echo.Context, status int, data interface{}) error {
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Result: data,
	})
}
