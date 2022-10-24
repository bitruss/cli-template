package api

import (
	"net/http"

	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/common/dbkv"
	"github.com/coreservice-io/cli-template/src/common/http/api"
	"github.com/coreservice-io/cli-template/src/user_mgr"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type DBKV struct {
	Id          int64  `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

//create
// @Description Msg_Req_CreateRule
type Msg_Req_CreateKv struct {
	Key         string `json:"key"`         //required
	Value       string `json:"value"`       //required
	Description string `json:"description"` //required
}

//query
// @Description Msg_Req_QueryKv_Filter
type Msg_Req_QueryKv_Filter struct {
	Keys *[]string `json:"keys"` //optional
}

// @Description Msg_Req_QueryKv
type Msg_Req_QueryKv struct {
	Filter Msg_Req_QueryKv_Filter `json:"filter"` //required
}

type Msg_Resp_QueryKv struct {
	api.API_META_STATUS
	Kv    []*DBKV `json:"kv"`
	Count int64   `json:"count"`
}

//delete
// @Description Msg_Req_DeleteKv_Filter
type Msg_Req_DeleteKv_Filter struct {
	Keys []string `json:"keys"` //required
}

// @Description Msg_Req_DeleteKv
type Msg_Req_DeleteKv struct {
	Filter Msg_Req_DeleteKv_Filter `json:"filter"` //required
}

//update
// @Description Msg_Req_UpdateKv_Filter
type Msg_Req_UpdateKv_Filter struct {
	Key string `json:"key"` //required
}

// @Description Msg_Req_UpdateKv_To
type Msg_Req_UpdateKv_To struct {
	Value       string `json:"value"`       //required
	Description string `json:"description"` //required
}

// @Description Msg_Req_UpdateKv
type Msg_Req_UpdateKv struct {
	Filter Msg_Req_UpdateKv_Filter `json:"filter"` //required
	Update Msg_Req_UpdateKv_To     `json:"update"` //required
}

func config_kv(httpServer *echo.Echo) {
	httpServer.POST("/api/kv/create", createKvHandler, MidToken)
	httpServer.POST("/api/kv/query", queryKvHandler, MidToken)
	httpServer.POST("/api/kv/delete", deleteKvHandler, MidToken)
	httpServer.POST("/api/kv/update", updateKvHandler, MidToken)
}

// @Summary      creat key value pair
// @Tags         kv
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  Msg_Req_CreateKv true  "creat key value pair"
// @Produce      json
// @Success      200 {object} api.API_META_STATUS "result"
// @Router       /api/kv/create [post]
func createKvHandler(ctx echo.Context) error {

	var msg Msg_Req_CreateKv
	res := &api.API_META_STATUS{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	userInfo, err := CheckUser(ctx)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}
	if !userInfo.HasOneOfRoles([]string{user_mgr.USER_ROLE_ADMIN}) {
		res.MetaStatus(-1, "auth error")
		return ctx.JSON(http.StatusOK, res)
	}

	kvResult, err := dbkv.QueryDBKV(sqldb_plugin.GetInstance(), nil, &[]string{msg.Key}, false, false)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	} else if len(kvResult.Kv) > 0 {
		res.MetaStatus(-1, "key already exist")
		return ctx.JSON(http.StatusOK, res)
	}

	err = dbkv.SetDBKV_Str(sqldb_plugin.GetInstance(), msg.Key, msg.Value, msg.Description)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}

	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}

// @Summary      query key value pair
// @Tags         kv
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  Msg_Req_QueryKv true  "query key value pair"
// @Produce      json
// @Success      200 {object} Msg_Resp_QueryKv "result"
// @Router       /api/kv/query [post]
func queryKvHandler(ctx echo.Context) error {

	var msg Msg_Req_QueryKv
	res := &Msg_Resp_QueryKv{}
	res.Kv = []*DBKV{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	fromCache := true
	if msg.Filter.Keys == nil {
		userInfo, err := CheckUser(ctx)
		if err != nil {
			res.MetaStatus(-1, err.Error())
			return ctx.JSON(http.StatusOK, res)
		}
		if !userInfo.HasOneOfRoles([]string{user_mgr.USER_ROLE_ADMIN, user_mgr.USER_ROLE_READONLY}) {
			if msg.Filter.Keys == nil {
				res.MetaStatus(-1, "auth error")
				return ctx.JSON(http.StatusOK, res)
			}
		} else {
			fromCache = false
		}
	}

	kvResult, err := dbkv.QueryDBKV(sqldb_plugin.GetInstance(), nil, msg.Filter.Keys, fromCache, true)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}

	copier.Copy(&res.Kv, &kvResult.Kv)
	res.Count = kvResult.TotalCount
	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}

// @Summary      delete key value pair
// @Tags         kv
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  Msg_Req_DeleteKv true  "delete key value pair"
// @Produce      json
// @Success      200 {object} api.API_META_STATUS "result"
// @Router       /api/kv/delete [post]
func deleteKvHandler(ctx echo.Context) error {

	var msg Msg_Req_DeleteKv
	res := &api.API_META_STATUS{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	userInfo, err := CheckUser(ctx)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}
	if !userInfo.HasOneOfRoles([]string{user_mgr.USER_ROLE_ADMIN}) {
		res.MetaStatus(-1, "auth error")
		return ctx.JSON(http.StatusOK, res)
	}

	if len(msg.Filter.Keys) == 0 {
		res.MetaStatus(-1, "key needed")
		return ctx.JSON(http.StatusOK, res)
	}
	if len(msg.Filter.Keys) > 1 {
		res.MetaStatus(-1, "only support one key")
		return ctx.JSON(http.StatusOK, res)
	}

	err = dbkv.DeleteDBKV_Key(sqldb_plugin.GetInstance(), msg.Filter.Keys[0])
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}

	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}

// @Summary      update key value pair
// @Tags         kv
// @Security     ApiKeyAuth
// @Accept       json
// @Param        msg  body  Msg_Req_UpdateKv true  "update key value pair"
// @Produce      json
// @Success      200 {object} api.API_META_STATUS "result"
// @Router       /api/kv/update [post]
func updateKvHandler(ctx echo.Context) error {

	var msg Msg_Req_UpdateKv
	res := &api.API_META_STATUS{}

	if err := ctx.Bind(&msg); err != nil {
		res.MetaStatus(-1, "bad request data")
		return ctx.JSON(http.StatusOK, res)
	}

	userInfo, err := CheckUser(ctx)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}
	if !userInfo.HasOneOfRoles([]string{user_mgr.USER_ROLE_ADMIN}) {
		res.MetaStatus(-1, "auth error")
		return ctx.JSON(http.StatusOK, res)
	}

	kvResult, err := dbkv.QueryDBKV(sqldb_plugin.GetInstance(), nil, &[]string{msg.Filter.Key}, false, false)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	} else if len(kvResult.Kv) == 0 {
		res.MetaStatus(-1, "key not exist")
		return ctx.JSON(http.StatusOK, res)
	}

	err = dbkv.SetDBKV_Str(sqldb_plugin.GetInstance(), msg.Filter.Key, msg.Update.Value, msg.Update.Description)
	if err != nil {
		res.MetaStatus(-1, err.Error())
		return ctx.JSON(http.StatusOK, res)
	}

	res.MetaStatus(1, "success")
	return ctx.JSON(http.StatusOK, res)
}
