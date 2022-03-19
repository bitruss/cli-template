package api

import (
	"strconv"

	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/src/examples/userMgr"
	"github.com/labstack/echo/v4"
)

func config_user(httpServer *echoServer.EchoServer) {
	//create
	httpServer.POST("/api/user/create", createUser, CheckToken)

	//get
	httpServer.GET("/api/user/get/:id", getUser, CheckToken)

	//update
	httpServer.POST("/api/user/update", updateUser, CheckToken)
}

// @Summary      creat user
// @Description  creat user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  CreateUserMsg true  "new user info"
// @Produce      json
// @Success      200 {object} echoServer.RespBody{result=userMgr.ExampleUserModel} "result"
// @Router       /api/user/create [post]
func createUser(ctx echo.Context) error {
	var msg CreateUserMsg
	if err := ctx.Bind(&msg); err != nil {
		return echoServer.ErrorResp(ctx, -1, nil, "post data error")
	}

	//todo create user in db
	//mock db action
	r := userMgr.ExampleUserModel{
		ID:     1,
		Status: "normal",
		Name:   msg.Name,
		Email:  msg.Email,
	}
	return echoServer.SuccessResp(ctx, 1, r, "")
}

// @Summary      get user
// @Description  get user
// @Tags         user
// @Security     ApiKeyAuth
// @Param        id  path  integer true  "user id"
// @Produce      json
// @Success      200 {object} echoServer.RespBody{result=userMgr.ExampleUserModel} "result"
// @Router       /api/user/get/{id} [get]
func getUser(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return echoServer.ErrorResp(ctx, -1, nil, "id error")
	}
	if id == 0 {
		return echoServer.ErrorResp(ctx, -1, nil, "id error")
	}

	//todo get user info from db
	//mock db action
	r := userMgr.ExampleUserModel{
		ID:      1,
		Status:  "normal",
		Name:    "jack",
		Email:   "jsck@email.com",
		Updated: 1647609300,
		Created: 1647609300,
	}

	return echoServer.SuccessResp(ctx, 1, r, "")
}

// @Summary      update user
// @Description  update user
// @Tags         user
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  UpdateUserMsg true  "update user info"
// @Produce      json
// @Success      200 {object} echoServer.RespBody "result"
// @Router       /api/user/update [post]
func updateUser(ctx echo.Context) error {
	var msg UpdateUserMsg
	if err := ctx.Bind(&msg); err != nil {
		return echoServer.ErrorResp(ctx, -1, nil, "post data error")
	}

	//update user
	updateData := map[string]interface{}{}
	if msg.ID == nil {
		return echoServer.ErrorResp(ctx, -1, nil, "user id is required")
	}
	if msg.Name != nil {
		updateData["name"] = *msg.Name
	}
	if msg.Email != nil {
		updateData["email"] = *msg.Email
	}
	if msg.Status != nil {
		updateData["status"] = *msg.Status
	}

	//todo update user info in db

	return echoServer.SuccessResp(ctx, 1, nil, "")
}

//example msg
type CreateUserMsg struct {
	Name  string
	Email string
}

type UpdateUserMsg struct {
	ID     *int
	Status *string
	Name   *string
	Email  *string
}
