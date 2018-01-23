package httpsvr

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Server ...
type Server struct {
	addr	string
	router	*httprouter.Router
	opt 	*option
	log 	log.Logger
	oriSvr	http.Server
}

func New(addr string, opts ...ServerOption) *Server {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}
	if addr == "" {
		addr = "127.0.0.1:10024"
	}
	s := &Server{
		addr: 	addr,
		router: httprouter.New(),
		opt:	opt,
	}
	s.oriSvr = &http.Server{Addr: addr, Handler: s}
	return s
}

// ServeHTTP implement net/http.router
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *Server) AddRoute(method, path string, ctrl IController){
	s.router.Handler()
}

func (s *Server) Server() error{
	return s.oriSvr.ListenAndServe()
}
