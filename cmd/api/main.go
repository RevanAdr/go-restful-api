package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RevanAdr/go-restful-api/internal/catalog"
	"github.com/RevanAdr/go-restful-api/internal/platform/config"
)

func main() {
	config := config.NewConfig()
	db := config.Load()

	defer db.Close()

	repo := catalog.NewCatalogRepository(db)
	service := catalog.NewCatalogService(repo)
	handler := catalog.NewCatalogHandler(service)
	router := catalog.NewCatalogRouter(handler)

	fmt.Println("Server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", router.Route()))
}
