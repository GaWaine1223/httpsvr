package httpsvr

import (
	"net/http"
	"fmt"
)

// GetClientAddr ...
func GetClientAddr(req *http.Request) string {
	if addr := req.Header.Get("HTTP_CLIENT_IP"); addr != "" {
		return addr
	} else if addr := req.Header.Get("HTTP_X_FORWARDED_FOR"); addr != "" {
		return addr
	}
	return req.RemoteAddr
}

func getErrMsg(err error) []byte {
	return []byte(fmt.Sprintf(`{"errno":-1,"errmsg":"%s"}`, err.Error()))
}