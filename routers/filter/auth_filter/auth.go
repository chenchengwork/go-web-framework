package auth_filter

import (
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/gin-gonic/gin"
)

type authParams struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

// 检验参数
func GetAuth(c *gin.Context) {
	err := c.ShouldBindQuery(&authParams{
		Username: c.Query("username"),
		Password: c.Query("password"),
	})

	if err != nil {
		app.MarkErrors(c, err)
		c.Abort()
		return
	}

	c.Next()
}
