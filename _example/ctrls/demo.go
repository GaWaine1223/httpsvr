// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package ctrls

import (
	"github.com/GaWaine1223/Lothar/httpsvr/_example/handler"
	"github.com/GaWaine1223/Lothar/httpsvr/_example/idl"
)

type DemoCtrl struct {
}

func (c *DemoCtrl) GenIdl() interface{} {
	return idl.NewDemoReq()
}

func (c *DemoCtrl) Do(req interface{}) interface{} {
	r := req.(*idl.DemoReq)
	return handler.DemoHandle(r)
}
