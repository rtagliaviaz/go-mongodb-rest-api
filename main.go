package main

import (
	"log"
	"net/http"

	"./db"
	"./router"
)

func main() {
	// Database
	db.Connect()

	r := router.Router()

	log.Fatal(http.ListenAndServe(":3000", r))
}
