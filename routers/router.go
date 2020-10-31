package routers

import (
	"github.com/chenchengwork/go-web-framework/middleware"
	"github.com/chenchengwork/go-web-framework/pkg/cookie"
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/chenchengwork/go-web-framework/routers/api"
	"github.com/chenchengwork/go-web-framework/routers/filter/auth_filter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 设置路由中间件
	r.Use(middleware.HttpsHandler())
	cookie.Setup(r) // 设置cookie

	// 设置静态资源访问目录
	r.StaticFS("/static", http.Dir(setting.AppSetting.RuntimeRootPath+"www"))

	// 配置访问api
	r.GET("/login", auth_filter.Login, api.Auth.Login)

	return r
}
