package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

// 反向代理中间件
func ReverseProxyHandler(host string, port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse("http://" + host + ":" + port)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
