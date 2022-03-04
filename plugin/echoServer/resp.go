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

//status <0
func ErrorResp(c echo.Context, status int, data interface{}, msg string) error {
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Result: data,
		Msg:    msg,
	})
}

//status >0
func SuccessResp(c echo.Context, status int, data interface{}, msg string) error {
	return c.JSON(http.StatusOK, RespBody{
		Status: status,
		Result: data,
		Msg:    msg,
	})
}
