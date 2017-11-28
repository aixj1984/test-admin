package payloads

import (
	"encoding/json"
	"flag"
	"test-admin/comm/utils"

	validator "gopkg.in/go-playground/validator.v9"
)

//Validate  validate payload 参数
var (
	validate *validator.Validate
	tagName  string
)

func init() {
	validate = validator.New()
	flag.StringVar(&tagName, "validatetag", "validate", " set log config file path")
}

//Init 初始化payload服务
func Init() {
	validate.SetTagName(tagName)
}

//GeneratePayload  生成payload
func GeneratePayload(data interface{}, body []byte) (err error) {
	err = json.Unmarshal(body, data)
	if err != nil {
		return utils.ErrParamsNotMatch
	}
	err = validate.Struct(data)
	if err == nil {
		return
	}
	return utils.ErrParamsNotMatch
}
