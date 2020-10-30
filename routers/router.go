package routers

import (
	"github.com/chenchengwork/go-web-framework/pkg/cookie"
	"github.com/chenchengwork/go-web-framework/routers/api"
	"github.com/chenchengwork/go-web-framework/routers/filter/auth_filter"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 设置路由中间件
	cookie.Setup(r) // 设置cookie

	// 设置静态资源访问目录

	// 配置访问api
	r.GET("/auth", auth_filter.GetAuth, api.Auth.GetAuth)

	return r
}
