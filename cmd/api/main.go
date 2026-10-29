package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RevanAdr/go-restful-api/internal/catalog"
	"github.com/RevanAdr/go-restful-api/internal/platform/config"
)

func main() {
	config := config.NewConfig()
	db, port := config.Load()

	formattedPort := fmt.Sprintf(":%s", port)

	defer db.Close()

	repo := catalog.NewCatalogRepository(db)
	service := catalog.NewCatalogService(repo)
	handler := catalog.NewCatalogHandler(service)
	router := catalog.NewCatalogRouter(handler)

	server := &http.Server{
		Addr:              formattedPort,
		Handler:           router.Route(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	fmt.Printf("Server listening on %s", formattedPort)
	log.Fatal(server.ListenAndServe())
}
