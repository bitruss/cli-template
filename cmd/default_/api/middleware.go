package api

import (
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/tools/http"
	"github.com/labstack/echo/v4"
)

func MidToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := http.GetBearToken(c.Request().Header)
		if token == "" {
			return echoServer.ErrorResp(c, -1, nil, "token error")
		}
		//get token related info and set it to context
		//for future role,permission management ,you can get it in api using ctx.Get
		//c.Set("tokenInfo",{your token struct})
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
