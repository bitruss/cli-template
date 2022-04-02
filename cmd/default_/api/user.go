package api

import (
	"net/http"

	"github.com/coreservice-io/CliAppTemplate/basic"
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
	Id     *[]int  //sql : id in (...) //optional
	Name   *string //optional
	Email  *string //optional  email can be like condition e.g " LIKE `%jack%` "
	Offset int     //required
	Limit  int     //required
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

	//pass qmap to your code inside your manager
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

type MSG_REQ_UPDATE_WHERE_USER struct {
	ID     *int
	Status *string
	Name   *string
	Email  *string
}

type MSG_REQ_UPDATE_TO_USER struct {
	Status *string
	Name   *string
	Email  *string
}

type MSG_REQ_UPDATE_USER struct {
	Where MSG_REQ_UPDATE_WHERE_USER
	To    MSG_REQ_UPDATE_TO_USER
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
// @Param        msg  body  MSG_REQ_UPDATE_USER true  "update user"
// @Produce      json
// @Success      200 {object} MSG_RESP_UPDATE_USER "result"
// @Router       /api/user/update [post]
func updateUser(ctx echo.Context) error {
	var msg MSG_REQ_UPDATE_USER
	var res MSG_RESP_UPDATE_USER
	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "post data error")
		return ctx.JSON(http.StatusOK, res)
	}

	qmap := data.MapRemoveNil(structs.Map(msg.Where))
	tomap := data.MapRemoveNil(structs.Map(msg.To))

	//pass qmap and tomap to your code inside your manager

	//do your work here
	basic.Logger.Debugln(qmap)
	basic.Logger.Debugln(tomap)

	//

	//todo update user info in db
	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}
