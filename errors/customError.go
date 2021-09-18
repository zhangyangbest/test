package errors

import "fmt"

type ErrCode int

//定义错误码
const (
	ERR_OK ErrCode = iota // 请求OK
	ERR_PARAMS
	ERR_TIMEOUT
)

// 定义错误码与描述信息的映射
var mapErrDesc = map[ErrCode]string {
	ERR_OK: "OK",
	ERR_PARAMS: "无效参数",
	ERR_TIMEOUT: "超时",
}

// 实现string() 这样print会调用string()
func (e ErrCode) String() string{
	if errDesc, ok := mapErrDesc[e]; ok {
		return errDesc
	} else {
		return fmt.Sprintf("error code %d", e)
	}
}
