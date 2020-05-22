package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Server struct
type Server struct {
	Router *mux.Router
}

//Handler struct
type Handler struct {
	Handler http.Handler
}

//Initialize : TODO PASS CONFIG
func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.SetRouters()
}

//SetRouters : set routers
func (s *Server) SetRouters() {
	h := &Handler{}
	s.Router.HandleFunc("/hello", h.HelloWorld).Methods(http.MethodGet)
}

//HelloWorld : return hello world message
func (h *Handler) HelloWorld(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("{\"respuesta\": \"Hola mundo\"}")
	return
}

//Run : start server
func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
