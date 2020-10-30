package api

import (
	"fmt"
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/chenchengwork/go-web-framework/pkg/cookie"
	"github.com/chenchengwork/go-web-framework/pkg/e"
	"github.com/chenchengwork/go-web-framework/pkg/util"
	"github.com/chenchengwork/go-web-framework/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authApi struct{}

var Auth authApi

type authParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func (api *authApi) GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.Query("username")
	password := c.Query("password")
	//err := c.ShouldBindQuery(&authParams{Username: username, Password: password})
	//
	//if err != nil{
	//	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS,app.MarkErrors(err))
	//
	//	return
	//}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	fmt.Println(cookie.Get(c))

	cookie.Save(c, &cookie.Value{
		UserId:   "xxxx",
		UserName: "chencheng",
	})

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
