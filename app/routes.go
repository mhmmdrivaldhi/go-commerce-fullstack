package app

import (
	"github.com/gorilla/mux"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/controllers"
)

func (server *Server) Routes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}