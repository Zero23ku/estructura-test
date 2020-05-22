package handler

import (
	"encoding/json"
	"net/http"
)

//Handler struct
type Handler struct {
	Handler http.Handler
}

//HelloWorld : return hello world message
func (h *Handler) HelloWorld(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("{\"respuesta\": \"Hola mundo\"}")
	return
}
