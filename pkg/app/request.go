package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chenchengwork/go-web-framework/pkg/e"
	"github.com/chenchengwork/go-web-framework/pkg/logging"
	"github.com/chenchengwork/go-web-framework/pkg/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func MarkErrors(c *gin.Context, err error) string {
	appG := Gin{C: c}
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

	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, errStr)

	return errStr
}
