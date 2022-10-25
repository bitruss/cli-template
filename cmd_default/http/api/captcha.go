package api

import (
	"net/http"

	"github.com/coreservice-io/cli-template/src/common/captcha"
	"github.com/coreservice-io/cli-template/src/common/http/api"
	"github.com/coreservice-io/cli-template/src/common/limiter"
	"github.com/labstack/echo/v4"
)

// @Msg_Resp_Captcha
type Msg_Resp_Captcha struct {
	api.API_META_STATUS
	Id      string `json:"id"`
	Content string `json:"content"`
}

func configCaptcha(httpServer *echo.Echo) {
	// user
	httpServer.GET("/api/captcha", getCaptchaHandler)
}

// @Summary      get captcha
// @Tags         captcha
// @Produce      json
// @response 	 200 {object} Msg_Resp_Captcha "result"
// @Router       /api/captcha [get]
func getCaptchaHandler(ctx echo.Context) error {

	res := &Msg_Resp_Captcha{}

	remoteIp := ctx.RealIP()

	// request rate limit
	if !limiter.Allow("captcha:"+remoteIp, 1, 3) {
		// error not cool down
		res.MetaStatus(-1, "captcha too frequently!")
		return ctx.JSON(http.StatusOK, res)
	}

	id, base64Code, err := captcha.GenCaptcha()
	if err != nil {
		// error gen captcha
		res.MetaStatus(-1, "gen captcha err")
		return ctx.JSON(http.StatusOK, res)
	}
	if id == "" || base64Code == "" {
		// error gen captcha
		res.MetaStatus(-1, "gen captcha err")
		return ctx.JSON(http.StatusOK, res)
	}

	res.MetaStatus(1, "success")
	res.Id = id
	res.Content = base64Code
	return ctx.JSON(http.StatusOK, res)
}
