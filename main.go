package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go-api/config"
	"example.com/go-api/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	// Start the HTTP server
	log.Println("Server started on : 3001")
	log.Fatal(http.ListenAndServe(":"+config.Ip_server, router.GetRouting()))
	godotenv.Load(".env")

}
