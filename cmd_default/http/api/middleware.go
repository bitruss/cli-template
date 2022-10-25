package api

import (
	"errors"

	"github.com/coreservice-io/cli-template/src/common/http"
	"github.com/coreservice-io/cli-template/src/user_mgr"
	"github.com/labstack/echo/v4"
)

func MidToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("token", http.GetBearToken(c.Request().Header))
		//continue
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func GetUserInfo(c echo.Context) *user_mgr.UserModel {
	v := c.Get("userInfo")
	if v == nil {
		return nil
	}
	return v.(*user_mgr.UserModel)
}

func CheckUser(c echo.Context) (*user_mgr.UserModel, error) {
	userInfo := GetUserInfo(c)
	if userInfo == nil {
		return nil, errors.New("user not exist")
	}
	if userInfo.Forbidden {
		return nil, errors.New("user forbidden")
	}
	return userInfo, nil
}
