package common

import "sync"

var (
	// okDataPool http响应成功数据池化
	okDataPool *sync.Pool

	// statusCodeMsg 错误代码消息
	statusCodeMsg map[StatusCode]string

	// emptyData 空数据
	emptyData = &struct{}{}
)

// ResData http响应数据封包
type ResData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Ok 返回成功响应数据封包
func Ok(data interface{}) *ResData {
	resData := okDataPool.Get().(*ResData)
	resData.Data = data
	return resData
}

// EmptyData 空数据
func EmptyData() *struct{} {
	return emptyData
}

// RecycleOk 回收Ok响应数据封包
func RecycleOk(data *ResData) {
	data.Data = nil // Notice:一定要赋nil避免泄内存
	okDataPool.Put(data)
}

// ErrCodeMsg 返回错误消息响应数据
func ErrCodeMsg(code StatusCode, msg ...string) *ResData {
	var errMsg string
	var exist bool
	if len(msg) > 0 {
		errMsg = msg[0]
	} else if errMsg, exist = statusCodeMsg[code]; !exist {
		errMsg = statusCodeMsg[CodeInternalServerError]
	}

	return &ResData{
		Code: int(code),
		Msg:  errMsg,
		Data: emptyData,
	}
}
