package httpsvr

import "time"

type Access struct {
	maxAccess int
	tk time.Ticker
	// 输入桶
	ibucket  chan struct{}
	// 输出桶
	obucket  chan struct{}
}

func NewAccessor(i int) *Access {
	var ac Access
	ac.maxAccess = i
	// 10ms 准入超时计时器，时间窗口
	ac.tk = time.NewTicker(time.Millisecond * 10)
	ac.ibucket = make(chan struct{}, i)
	ac.obucket = make(chan struct{}, i)
	return ac
}

func (a *Access) Run() {

}

func (a *Access) InControl(ic chan struct{}) error {
	select {
	case ic <- struct{}(1) :
	case <- a.tk.C :

	}
}

func (a *Access) OutControl(oc chan struct{}) {

}
