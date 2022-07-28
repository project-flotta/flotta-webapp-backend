package api

import (
	"github.com/ahmadateya/flotta-webapp-backend/api/handlers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HelloServer)

	return r
}
