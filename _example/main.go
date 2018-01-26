package	main

import (
	"time"
	"github.com/GaWaine1223/Lothar/httpsvr"
	"github.com/GaWaine1223/Lothar/httpsvr/_example/ctrls"
)

func main() {
	s := httpsvr.New("127.0.0.1:10024",
		httpsvr.SetReadTimeout(time.Millisecond*200),
		httpsvr.SetWriteTimeout(time.Millisecond*200),
		httpsvr.SetMaxAccess(100),
	)
	s.AddRoute("POST", "/test/api", &ctrls.DemoCtrl{})
	s.Serve()
}
