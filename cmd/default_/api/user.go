package api

import (
	"net/http"

	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/tools/data"
	"github.com/coreservice-io/CliAppTemplate/tools/http/api"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

func config_user(httpServer *echoServer.EchoServer) {
	//create
	httpServer.POST("/api/user/create", createUser, MidToken)

	//get
	httpServer.GET("/api/user/search", searchUser, MidToken)

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

type MSG_REQ_SEARCH_USER struct {
	api.API_META_STATUS
	Id         *[]int  //sql : id in (...) //optional
	Name       *string //optional
	Email_like *string //optional
}

type MSG_USER struct {
	Id    int
	Name  string
	Email string
}

type MSG_RESP_SEARCH_USER struct {
	api.API_META_STATUS
	Result []*MSG_USER
}

// @Summary      search user
// @Description  search user
// @Tags         user
// @Security     ApiKeyAuth
// @Param        msg  body  MSG_REQ_SEARCH_USER true  "user search param"
// @Produce      json
// @Success      200 {object} MSG_RESP_GET_USER "result"
// @Router       /api/user/search [post]
func searchUser(ctx echo.Context) error {

	var msg MSG_REQ_SEARCH_USER
	res := &MSG_RESP_SEARCH_USER{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	qmap := data.MapRemoveNil(structs.Map(msg))

	//do this part in your manager code
	if len(qmap) == 0 {
		res.MetaStatus(-1, "no query condition ")
		return ctx.JSON(http.StatusOK, res)
	}

	//fill your res ,mock db action

	//end of manager code

	//todo get user info from db
	res.MetaStatus(1, "success")
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
