package api

import (
	"time"

	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/labstack/echo/v4"
)

// @Summary      health check
// @Description  health check
// @Tags         other
// @Produce      plain
// @Success      200 {object} echoServer.RespBody{data=int64} "result"
// @Router       /api/health [get]
func healthHandler(ctx echo.Context) error {
	return echoServer.SuccessResp(ctx, 1, time.Now().Unix(), "")
}