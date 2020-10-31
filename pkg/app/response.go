package app

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

//type Gin struct {
//	C *gin.Context
//}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
//func (g *Gin) Response(httpCode, errCode int, data interface{}) {
//	g.C.JSON(httpCode, Response{
//		Code: errCode,
//		Msg:  e.GetMsg(errCode),
//		Data: data,
//	})
//	return
//}

func ResponseSuccess(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
	return
}

// 系统未知错误http code 500
func ResponseFailed(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, response{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: data,
	})
	return
}

// 未登录http code 401
func ResponseAuthFailed(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusUnauthorized, response{
		Code: http.StatusUnauthorized,
		Msg:  msg,
		Data: data,
	})
	return
}

// 校验参数未通过http code 400
func ResponseCheckParamsFailed(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, response{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: data,
	})
	return
}

// 未找到资源http code 404
func ResponseNoFound(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusNotFound, response{
		Code: http.StatusNotFound,
		Msg:  msg,
		Data: data,
	})
	return
}

// 未受权限http code 403
func ResponseNoPermission(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusForbidden, response{
		Code: http.StatusForbidden,
		Msg:  msg,
		Data: data,
	})
	return
}
