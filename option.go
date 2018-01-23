package httpsvr

import "time"

type option struct {
	dumpResponse  bool
	enableElasped bool
	dumpAccess    bool
	validate      bool
	readTimeout   time.Duration
	writeTimeout  time.Duration
}

type ServerOption func(o *option)

func SetReadTimeOut(rt time.Duration) ServerOption {
	return func(o *option) {
		o.readTimeout = rt
	}
}

func SetWriteTimeOut(wt time.Duration) ServerOption {
	return func(o *option) {
		o.writeTimeout = wt
	}
}
