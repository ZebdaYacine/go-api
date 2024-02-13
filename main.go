package main

import (
	"log"
	"net/http"

	"example.com/go-api/router"
)

func main() {
	// Start the HTTP server
	log.Println("Server started on : 3001")
	log.Fatal(http.ListenAndServe(":3001", router.GetRouting()))

}
