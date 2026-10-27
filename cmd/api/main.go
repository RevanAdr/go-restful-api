package main

import (
	"log"
	"net/http"

	"github.com/RevanAdr/go-restful-api/internal/handler"
)

func main() {
	router := handler.NewRouter()

	log.Println("Server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
