package httpsvr

//Response 响应结构体
type Response struct {
	Code int         `json:"errno"`
	Msg  string      `json:"errmsg"`
	Data interface{} `json:"data"`
}
