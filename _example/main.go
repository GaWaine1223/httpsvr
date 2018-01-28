// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GaWaine1223/Lothar/httpsvr"
	"github.com/GaWaine1223/Lothar/httpsvr/_example/ctrls"
)

func main() {
	s := httpsvr.New("127.0.0.1:10024",
		httpsvr.SetReadTimeout(time.Millisecond*200),
		httpsvr.SetWriteTimeout(time.Millisecond*200),
		httpsvr.SetMaxAccess(2),
	)
	go GracefulExit(s)
	s.AddRoute("POST", "/test/api", &ctrls.DemoCtrl{})
	s.Serve()
}

// GracefulExit 优雅退出
func GracefulExit(svr *httpsvr.Server) {
	sigc := make(chan os.Signal, 0)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	<-sigc
	println("closing agent...")
	svr.GracefulExit()
	println("agent closed.")
	os.Exit(0)
}
