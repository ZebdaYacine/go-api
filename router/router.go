package router

import (
	Handler "example.com/go-api/handler"
	"github.com/gorilla/mux"
)

func GetRouting() *mux.Router {
	router := mux.NewRouter()
	// Define routes
	router.HandleFunc("/", Handler.HomeHandler).Methods("GET")
	router.HandleFunc("/about", Handler.AboutHandler).Methods("GET")
	return router
}
