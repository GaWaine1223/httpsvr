package httpsvr

// IController ...
type IController interface {
	GenIdl() interface{}
	Do(interface{}) interface{}
}
