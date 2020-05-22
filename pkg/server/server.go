package server

import (
	"log"
	"net/http"

	"github.com/Zero23ku/estructura-test/pkg/handler"
	"github.com/gorilla/mux"
)

//Server struct
type Server struct {
	Router *mux.Router
}

//Initialize : TODO PASS CONFIG
func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.SetRouters()
}

//SetRouters : set routers
func (s *Server) SetRouters() {
	h := &handler.Handler{}
	s.Router.HandleFunc("/hello", h.HelloWorld).Methods(http.MethodGet)
}

//Run : start server
func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
