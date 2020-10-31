package cookie

import (
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Value struct {
	UserId   string
	UserName string
}

func Setup(r *gin.Engine) {
	store := cookie.NewStore([]byte(setting.CookieSetting.Name))
	r.Use(sessions.Sessions(setting.CookieSetting.Name, store))
}
