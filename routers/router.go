package routers

import (
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/chenchengwork/go-web-framework/routers/api"
	v1 "github.com/chenchengwork/go-web-framework/routers/api/v1"
	"github.com/chenchengwork/go-web-framework/routers/filter/auth_filter"
	"github.com/chenchengwork/go-web-framework/routers/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 设置路由中间件
	r.Use(middleware.ErrorHandler(os.Stdout))
	r.Use(middleware.HttpsHandler())
	r.Use(middleware.CookieHandler())

	// 设置静态资源访问目录
	r.StaticFS("/static", http.Dir(setting.AppSetting.RuntimeRootPath+"www"))

	// 配置访问api
	r.GET("/login", auth_filter.Login, api.Auth.Login)

	// 以下为需要认证登录的api
	apiv1 := r.Group("/api/v1", middleware.AuthHandler())
	{
		apiv1.GET("/getUser", v1.User.GetUser)
	}

	return r
}
