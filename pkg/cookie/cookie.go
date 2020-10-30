package cookie

import (
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var defaultKey string = "Jwt"

type Value struct {
	UserId   string
	UserName string
}

func Setup(r *gin.Engine) {
	store := cookie.NewStore([]byte(setting.CookieSetting.Name))
	r.Use(sessions.Sessions(setting.CookieSetting.Name, store))
}

func Save(c *gin.Context, value *Value) {
	session := sessions.Default(c)
	session.Set("UserId", value.UserId)
	session.Set("UserName", value.UserName)
	session.Options(sessions.Options{MaxAge: setting.CookieSetting.MaxAge})
	session.Save()
}

func Get(c *gin.Context) *Value {
	session := sessions.Default(c)
	var value = Value{}
	value.UserId = session.Get("UserId").(string)
	value.UserName = session.Get("UserName").(string)

	return &value
}

func Del(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: -1})
	session.Clear()
	session.Save()
}
