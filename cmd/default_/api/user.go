package api

import (
	"net/http"
	"strconv"

	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/tools/http/api"
	"github.com/labstack/echo/v4"
)

func config_user(httpServer *echoServer.EchoServer) {
	//create
	httpServer.POST("/api/user/create", createUser, MidToken)

	//get
	httpServer.GET("/api/user/get/:id", getUser, MidToken)

	//update
	httpServer.POST("/api/user/update", updateUser, MidToken)
}

type MSG_REQ_CREATE_USER struct {
	Name  string
	Email string
}

type MSG_RESP_CREATE_USER struct {
	api.API_META_STATUS
	Name  string `json:"name"`
	Email string `json:"email"`
}

// @Summary      creat user
// @Description  creat user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  MSG_REQ_CREATE_USER true  "new user info"
// @Produce      json
// @Success      200 {object} MSG_RESP_CREATE_USER "result"
// @Router       /api/user/create [post]
func createUser(ctx echo.Context) error {

	var msg MSG_REQ_CREATE_USER
	res := &MSG_RESP_CREATE_USER{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}
	//todo create user in db
	//mock db action
	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}

type MSG_RESP_GET_USER struct {
	api.API_META_STATUS
	Name  string
	Email string
}

// @Summary      get user
// @Description  get user
// @Tags         user
// @Security     ApiKeyAuth
// @Param        id  path  integer true  "user id"
// @Produce      json
// @Success      200 {object} MSG_RESP_GET_USER "result"
// @Router       /api/user/get/{id} [get]
func getUser(ctx echo.Context) error {

	res := &MSG_RESP_GET_USER{}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res.MetaStatus(-1, "id error")
		return ctx.JSON(http.StatusOK, res)
	}
	if id == 0 {
		res.MetaStatus(-1, "id 0 bad ")
		return ctx.JSON(http.StatusOK, res)
	}

	//todo get user info from db
	//mock db action
	res.MetaStatus(1, "success")
	res.Name = "jack"
	res.Email = "jsck@email.com"
	return ctx.JSON(http.StatusOK, res)
}

type MSG_REQ_UPDATE_USER struct {
	ID     *int
	Status *string
	Name   *string
	Email  *string
}

type MSG_RESP_UPDATE_USER struct {
	api.API_META_STATUS
	Name  string `json:"name"`
	Email string `json:"email"`
}

// @Summary      update user
// @Description  update user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  MSG_REQ_UPDATE_USER true  "update user info"
// @Produce      json
// @Success      200 {object} MSG_RESP_UPDATE_USER "result"
// @Router       /api/user/update [post]
func updateUser(ctx echo.Context) error {
	var res MSG_RESP_UPDATE_USER
	var msg MSG_REQ_UPDATE_USER
	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "post data error")
		return ctx.JSON(http.StatusOK, res)
	}

	//update user
	if msg.ID == nil {
		res.MetaStatus(-1, "user id is required")
		return ctx.JSON(http.StatusOK, res)
	}
	//mock update
	//todo update user info in db
	res.MetaStatus(1, "success")
	res.Name = "mock_update_name"
	res.Email = "mock_update_email"

	return ctx.JSON(http.StatusOK, res)
}
