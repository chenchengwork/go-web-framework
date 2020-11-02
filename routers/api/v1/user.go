package v1

import (
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/gin-gonic/gin"
)

type userApi struct{}

var User userApi

// @Summary user
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func (api *userApi) GetUser(c *gin.Context) {

	app.ResponseSuccess(c, "获取用户列表成功", nil)
}
