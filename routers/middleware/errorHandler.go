package middleware

import (
	"fmt"
	"github.com/chenchengwork/go-web-framework/pkg/app"
	"github.com/chenchengwork/go-web-framework/pkg/logging"
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http/httputil"
	"runtime"
)

func ErrorHandler(out io.Writer) gin.HandlerFunc {
	logger := log.New(out, "\n\n\x1b[31m", log.LstdFlags)

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				//打印堆栈
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

				if setting.ServerSetting.RunMode == "debug" {
					logger.Printf("[Recovery] panic recovered:\n%s\n%s\n%s", string(httpRequest), err, stack)
				} else {
					logging.Error(fmt.Sprintf("[Recovery] panic recovered:\n%s\n%s\n%s", string(httpRequest), err, stack))
				}

				//返回统一格式数据
				app.ResponseFailed(ctx, "服务报错!", nil)
			}
		}()

		ctx.Next()
	}
}
