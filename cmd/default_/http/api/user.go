package api

import (
	"net/http"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	"github.com/coreservice-io/cli-template/tools/data"
	"github.com/coreservice-io/cli-template/tools/http/api"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

func config_user(httpServer *echo_plugin.EchoServer) {
	//create
	httpServer.POST("/api/user/create", createUser, MidToken)

	//get
	httpServer.GET("/api/user/search", searchUser, MidToken)

	//update
	httpServer.POST("/api/user/update", updateUser, MidToken)
}

type MSG_User struct {
	Id    int
	Name  string
	Email string
}

//create
type MSG_REQ_CREATE_USER struct {
	Name  string
	Email string
}

type MSG_RESP_CREATE_USER struct {
	api.API_META_STATUS
	User *MSG_User
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

//search
type MSG_REQ_SearchUser_Filter struct {
	Id    *[]int  //sql : id in (...) //optional
	Name  *string //optional
	Email *string //optional  email can be like condition e.g " LIKE `%jack%` "
}

type MSG_REQ_SearchUser struct {
	Filter MSG_REQ_SearchUser_Filter
	Offset int //required
	Limit  int //required
}

type MSG_RESP_SearchUser struct {
	api.API_META_STATUS
	Result []*MSG_User
}

// @Summary      search user
// @Description  search user
// @Tags         user
// @Security     ApiKeyAuth
// @Param        msg  body  MSG_REQ_SearchUser true  "user search param"
// @Produce      json
// @Success      200 {object} MSG_RESP_SearchUser "result"
// @Router       /api/user/search [post]
func searchUser(ctx echo.Context) error {

	var msg MSG_REQ_SearchUser
	res := &MSG_RESP_SearchUser{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	qmap := data.MapRemoveNil(structs.Map(msg.Filter))
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

type MSG_REQ_UpdateUser_Filter struct {
	ID []int
}

type Msg_Req_UpdateUser_To struct {
	Status *string
	Name   *string
	Email  *string
}

type MSG_REQ_UpdateUser struct {
	Filter MSG_REQ_UpdateUser_Filter
	Update Msg_Req_UpdateUser_To
}

// @Summary      update user
// @Description  update user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  MSG_REQ_UpdateUser true  "update user"
// @Produce      json
// @Success      200 {object} api.API_META_STATUS "result"
// @Router       /api/user/update [post]
func updateUser(ctx echo.Context) error {
	var msg MSG_REQ_UpdateUser
	var res api.API_META_STATUS
	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "post data error")
		return ctx.JSON(http.StatusOK, res)
	}

	qmap := data.MapRemoveNil(structs.Map(msg.Filter))
	tomap := data.MapRemoveNil(structs.Map(msg.Update))

	//pass qmap and tomap to your code inside your manager

	//do your work here
	basic.Logger.Debugln(qmap)
	basic.Logger.Debugln(tomap)

	//

	//todo update user info in db
	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}
