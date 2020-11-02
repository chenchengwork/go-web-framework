package middleware

import (
	"github.com/chenchengwork/go-web-framework/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func HttpsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			//SSLRedirect: true,
			//SSLHost:     "localhost:8000",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			logging.Error(err)
			c.Abort()
			return
		}

		c.Next()
	}
}
