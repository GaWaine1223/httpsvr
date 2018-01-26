package httpsvr

type IController interface {
	GenIdl() interface{}
	Do(interface{}) (interface{})
}
