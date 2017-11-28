package utils

import "errors"

//errors
var (
	ErrSuccess   = errors.New("成功")
	ErrNotUpdate = errors.New("版本无须更新")

	ErrParamsNotMatch = errors.New("参数缺失或类型不匹配")

	ErrIMEIDisabled    = errors.New("IMEI已经禁用")
	ErrGroupDisabled   = errors.New("分组已禁用")
	ErrProductDisabled = errors.New("产品已禁用")
	ErrEditionDisabled = errors.New("版本已禁用")
	ErrPackageDisabled = errors.New("升级包已禁用")
	ErrEditionNotFound = errors.New("版本不存在")
	ErrPackageNotFound = errors.New("升级包不存在")
	ErrProductNotFound = errors.New("产品不存在")

	ErrServerError = errors.New("服务器错误")
	ErrNotFound    = errors.New("接口不存在")
	ErrSign        = errors.New("SIGN_ERROR")

	ErrDataBase     = errors.New("数据库错误")
	ErrVersionRules = errors.New("版本规则出错")
)

var (
	errMessage = map[error]string{
		ErrSuccess:        "成功",
		ErrNotUpdate:      "版本无须更新",
		ErrParamsNotMatch: "参数缺失或类型不匹配",

		ErrIMEIDisabled:    "IMEI已经禁用",
		ErrGroupDisabled:   "分组已禁用",
		ErrProductDisabled: "产品已禁用",
		ErrEditionDisabled: "版本已禁用",
		ErrPackageDisabled: "升级包已禁用",
		ErrEditionNotFound: "版本不存在",
		ErrPackageNotFound: "升级包不存在",
		ErrProductNotFound: "产品不存在",

		ErrServerError: "服务器错误",
		ErrNotFound:    "接口不存在",

		ErrDataBase:     "数据库错误",
		ErrVersionRules: "版本规则出错",
		ErrSign:         "签名错误",
	}
)
var (
	errCode = map[error]int{
		ErrSuccess:      0,
		ErrNotUpdate:    100,
		ErrVersionRules: 101,

		ErrIMEIDisabled:    201,
		ErrGroupDisabled:   201,
		ErrProductDisabled: 201,
		ErrEditionDisabled: 201,
		ErrEditionNotFound: 201,
		ErrPackageDisabled: 201,
		ErrPackageNotFound: 201,
		ErrProductNotFound: 201,

		ErrServerError: 500,
		ErrNotFound:    404,

		ErrDataBase:       601,
		ErrParamsNotMatch: 602,
		ErrSign:           603,
	}
)

//Response  请求回复
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

//NewResponse response data
func NewResponse(err error, data ...interface{}) *Response {
	code, ok := errCode[err]
	if !ok {
		err = ErrServerError
		code = errCode[err]
	}
	message, _ := errMessage[err]
	res := &Response{
		Code:    code,
		Message: message,
	}
	if len(data) > 0 && data[0] != nil {
		res.Data = data[0]
	}
	return res
}
