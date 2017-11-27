package utils

type Resp struct {
	Succ int
	Msg  string
	Data interface{}
}

// RespStruct is the common function for response struction
// `succ` is an integer, `1` represents success, `0` represents failed.
// `msg` is an empty string when `succ` is 1, when `succ` is 0, `msg` will be the reason why it failed
// `data` is the response data
func RespStruct(succ int, msg string, data interface{}) Resp {
	var resp Resp
	resp.Succ = succ
	resp.Msg = msg
	resp.Data = data

	return resp
}
