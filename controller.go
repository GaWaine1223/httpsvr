// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package httpsvr

// IController ...
type IController interface {
	GenIdl() interface{}
	Do(interface{}) interface{}
}
