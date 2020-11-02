package middleware

import (
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CookieHandler() gin.HandlerFunc {
	store := cookie.NewStore([]byte(setting.CookieSetting.Name))

	return sessions.Sessions(setting.CookieSetting.Name, store)
}
