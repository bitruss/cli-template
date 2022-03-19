package api

import (
	"time"

	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/labstack/echo/v4"
)

// @Summary      health check
// @Description  health check
// @Tags         health
// @Produce      json
// @Success      200 {object} echoServer.RespBody{data=int64} "result"
// @Router       /api/health [get]
func healthCheck(ctx echo.Context) error {
	return echoServer.SuccessResp(ctx, 1, time.Now().Unix(), "")
}

func config_health(httpServer *echoServer.EchoServer) {
	//health
	httpServer.GET("/api/health", healthCheck)
}
