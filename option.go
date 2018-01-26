package httpsvr

import "time"

type option struct {
	maxAccess     int
	dumpResponse  bool
	enableElasped bool
	dumpAccess    bool
	validate      bool
	readTimeout   time.Duration
	writeTimeout  time.Duration
}

type ServerOption func(o *option)

func SetReadTimeout(rt time.Duration) ServerOption {
	return func(o *option) {
		o.readTimeout = rt
	}
}

func SetWriteTimeout(wt time.Duration) ServerOption {
	return func(o *option) {
		o.writeTimeout = wt
	}
}

func SetMaxAccess(i int) ServerOption {
	return func(o *option) {
		o.maxAccess = i
	}
}
