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
	router.HandleFunc("/search", Handler.SearchHandler).Methods("POST")
	router.HandleFunc("/Ask", Handler.AskQuestion).Methods("POST")
	return router
}
