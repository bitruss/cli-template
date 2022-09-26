package api

import (
	"net/http"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/src/common/data"
	"github.com/coreservice-io/cli-template/src/common/http/api"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

func config_user(httpServer *echo.Echo) {
	//create
	httpServer.POST("/api/user/create", createUser, MidToken)

	//get
	httpServer.GET("/api/user/search", searchUser, MidToken)

	//update
	httpServer.POST("/api/user/update", updateUser, MidToken)
}

type MSG_user struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create
type MSG_REQ_create_user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MSG_RESP_create_user struct {
	api.API_META_STATUS
	User *MSG_user `json:"user"`
}

// @Summary      creat user
// @Description  creat user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  MSG_REQ_create_user true  "new user info"
// @Produce      json
// @Success      200 {object} MSG_RESP_create_user "result"
// @Router       /api/user/create [post]
func createUser(ctx echo.Context) error {

	var msg MSG_REQ_create_user
	res := &MSG_RESP_create_user{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}
	//todo create user in db
	//mock db action
	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}

// search
type MSG_REQ_search_user_Filter struct {
	Id    *[]int  `json:"id"`    //sql : id in (...) //optional
	Name  *string `json:"name"`  //optional
	Email *string `json:"email"` //optional  email can be like condition e.g " LIKE `%jack%` "
}

type MSG_REQ_search_user struct {
	Filter MSG_REQ_search_user_Filter
	Offset int `json:"offset"` //required
	Limit  int `json:"limit"`  //required
}

type MSG_RESP_search_user struct {
	api.API_META_STATUS
	Result []*MSG_user `json:"result"`
}

// @Summary      search user
// @Description  search user
// @Tags         user
// @Security     ApiKeyAuth
// @Param        msg  body  MSG_REQ_search_user true  "user search param"
// @Produce      json
// @Success      200 {object} MSG_RESP_search_user "result"
// @Router       /api/user/search [post]
func searchUser(ctx echo.Context) error {

	var msg MSG_REQ_search_user
	res := &MSG_RESP_search_user{}

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

type MSG_REQ_update_user_Filter struct {
	ID []int `json:"id"`
}

type Msg_Req_update_user_to struct {
	Status *string `json:"status"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
}

type MSG_REQ_update_user struct {
	Filter MSG_REQ_update_user_Filter `json:"filter"`
	Update Msg_Req_update_user_to     `json:"update"`
}

// @Summary      update user
// @Description  update user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  MSG_REQ_update_user true  "update user"
// @Produce      json
// @Success      200 {object} api.API_META_STATUS "result"
// @Router       /api/user/update [post]
func updateUser(ctx echo.Context) error {
	var msg MSG_REQ_update_user
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
