package middleware

import (
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		// 校验是否登录
		session := sessions.Default(c)
		loginUserId := session.Get("UserId")
		if loginUserId == nil {
			app.ResponseAuthFailed(c, "请先登录", nil)
			c.Abort()
			return
		} else {
			session.Options(sessions.Options{
				MaxAge:   setting.CookieSetting.MaxAge,
				Path:     setting.CookieSetting.Path,
				HttpOnly: setting.CookieSetting.HttpOnly,
				Secure:   setting.CookieSetting.Secure,
				SameSite: setting.CookieSetting.GetSameSite(),
			})
			session.Save()
		}

		c.Next()
	}
}
