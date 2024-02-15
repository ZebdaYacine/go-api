package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Start the HTTP server
	log.Println(">>>>>>>>>>>" + os.Getenv("DB_HOST"))
	log.Println("Server started on : 3001")
	//log.Fatal(http.ListenAndServe(":3001", router.GetRouting()))
	godotenv.Load(".env")

}
