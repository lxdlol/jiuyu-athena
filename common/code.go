package common

import (
	"net/http"
	"sync"
)

// StatusCode 状态码
type StatusCode int

const (
	// CodeUnknow 未知
	CodeUnknow StatusCode = -1

	// CodeOk 请求响应
	CodeOk StatusCode = http.StatusOK

	// CodeAuth 暂无权限
	CodeAuth StatusCode = 403

	// CodeInternalServerError 内部服务出错
	CodeInternalServerError StatusCode = http.StatusInternalServerError

	// CodeIllegalParam 参数不合法
	CodeIllegalParam StatusCode = 400

	// CodeAuthLogin 未授权需登录
	CodeAuthLogin StatusCode = http.StatusUnauthorized

	// CodeIllegalPwdLength 请求参数必需要有token
	CodeIllegalPwdLength StatusCode = 10001
	//
	CodeUserError StatusCode = 10002
	//
	CodeRCodeError StatusCode = 10003
	//用户名已存在
	CodeIsExist StatusCode = 10004
)

// Error 实现error接口
func (c StatusCode) Error() string {
	if msg, ok := statusCodeMsg[c]; ok {
		return msg
	}
	return statusCodeMsg[CodeUnknow]
}

// init init
func init() {
	okDataPool = &sync.Pool{
		New: func() interface{} {
			return &ResData{
				Code: int(CodeOk),
				Msg:  "ok",
			}
		},
	}

	statusCodeMsg = map[StatusCode]string{
		CodeUnknow:              "unknow status",
		CodeOk:                  "操作成功",
		CodeInternalServerError: "服务繁忙,请稍后！",
		CodeIllegalParam:        "数据格式不正确",
		CodeAuthLogin:           "未授权，需登录",
		CodeAuth:                "没有访问权限",
		CodeIllegalPwdLength:    "密码长度错误",
		CodeUserError:           "用户名错误",
		CodeIsExist:             "用户名已存在",
	}
}
