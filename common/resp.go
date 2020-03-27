package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Resp 返回结构
type Resp struct{}

// Success 获取数据成功后的json数据响应
func (r *Resp) Success(c *gin.Context, data interface{}) {
	okData := Ok(data)
	defer RecycleOk(okData)
	c.JSON(http.StatusOK, okData)
	return
}

// Failure 获取数据失败后的json数据响应
func (r *Resp) Failure(c *gin.Context, errs ...error) {
	var (
		data   *ResData
		code   StatusCode
		errMsg string
	)

	err := errs[0]
	if len(errs) > 1 {
		errMsg = errs[1].Error()
	}

	if ok := errors.As(err, &code); ok {
		data = ErrCodeMsg(code)
	} else {
		data = ErrCodeMsg(CodeInternalServerError)
		if code, _ := strconv.Atoi(err.Error()); code > 0 {
			data = ErrCodeMsg(StatusCode(code))
		}
	}

	logrus.Errorf("request url: %s, failure: %v, err: %v", c.Request.URL, *data, errMsg)
	c.JSON(http.StatusOK, data)
	return
}
