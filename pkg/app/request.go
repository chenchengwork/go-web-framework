package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chenchengwork/go-web-framework/pkg/logging"
	"github.com/chenchengwork/go-web-framework/pkg/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func MarkErrors(c *gin.Context, err error) string {
	var errStr string

	switch err.(type) {
	case validator.ValidationErrors:
		errStr = validation.Translate(err.(validator.ValidationErrors))
	case *json.UnmarshalTypeError:
		unmarshalTypeError := err.(*json.UnmarshalTypeError)
		errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
	default:
		errStr = errors.New("unknown error.").Error()
	}

	logging.Error(err)

	ResponseAuthFailed(c, "参数校验失败", errStr)

	return errStr
}
