package ctrls

import (
	"github.com/GaWaine1223/Lothar/httpsvr/_example/idl"
	"github.com/GaWaine1223/Lothar/httpsvr/_example/handler"
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




