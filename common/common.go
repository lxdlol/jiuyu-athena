package common

var (
	// GResp 返回
	GResp *Resp
	// EmptyStruct 空结构体
	EmptyStruct = struct{}{}
)

func init() {
	// GResp 初始化
	GResp = new(Resp)
}
