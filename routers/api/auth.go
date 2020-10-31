package api

import (
	"fmt"
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/chenchengwork/go-web-framework/service/auth_service"
	"github.com/gin-gonic/gin"
)

type authApi struct{}

var Auth authApi

// @Summary login
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func (api *authApi) Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(password)

	//-------------------------------------------------------------------------
	// 此处添加校验登录的代码
	//-------------------------------------------------------------------------
	fmt.Println(auth_service.GetLoginUser(c))
	auth_service.SaveLoginUser(c, auth_service.LoginUser{
		UserId:    "xxx",
		AccountId: "accountId",
		UserName:  username,
		UserEmail: "chencheng@163.com",
		UserPhone: "150xxxxxxxx",
	})

	app.ResponseSuccess(c, "登录成功", nil)
}

func (api *authApi) Logout(c *gin.Context) {
	auth_service.Logout(c)

	app.ResponseSuccess(c, "退出成功", nil)
}
