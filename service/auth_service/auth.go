package auth_service

import (
	"github.com/chenchengwork/go-web-framework/models"
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

type LoginUser struct {
	Id        int64
	UserId    string
	AccountId string
	UserName  string
	UserEmail string
	UserPhone string
}

func SaveLoginUser(c *gin.Context, loginUserInfo LoginUser) {
	session := sessions.Default(c)

	session.Set("UserId", loginUserInfo.UserId)
	session.Set("AccountId", loginUserInfo.AccountId)
	session.Set("UserEmail", loginUserInfo.UserEmail)
	session.Set("UserPhone", loginUserInfo.UserPhone)
	session.Set("UserName", loginUserInfo.UserName)

	session.Options(sessions.Options{
		MaxAge:   setting.CookieSetting.MaxAge,
		Path:     setting.CookieSetting.Path,
		HttpOnly: setting.CookieSetting.HttpOnly,
		Secure:   setting.CookieSetting.Secure,
		SameSite: setting.CookieSetting.GetSameSite(),
	})

	session.Save()
}

func GetLoginUser(c *gin.Context) *LoginUser {
	session := sessions.Default(c)
	value := &LoginUser{}
	if session.Get("UserId") != nil {
		value.UserId = session.Get("UserId").(string)
	}

	if session.Get("AccountId") != nil {
		value.AccountId = session.Get("AccountId").(string)
	}
	if session.Get("UserEmail") != nil {
		value.UserEmail = session.Get("UserEmail").(string)
	}
	if session.Get("UserPhone") != nil {
		value.UserPhone = session.Get("UserPhone").(string)
	}
	if session.Get("UserName") != nil {
		value.UserName = session.Get("UserName").(string)
	}

	return value
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: -1})
	session.Clear()
	session.Save()
}
