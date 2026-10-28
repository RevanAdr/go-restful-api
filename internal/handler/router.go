package handler

import (
	"net/http"

	"github.com/RevanAdr/go-restful-api/internal/config"
	"github.com/RevanAdr/go-restful-api/internal/repository"
	"github.com/RevanAdr/go-restful-api/internal/service"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	db := config.Load()
	repo := repository.NewCategoryRepository()
	service := service.NewCategoryService(repo, db)
	handler := NewHandler(service)

	mux.HandleFunc("GET /api/v1/categories/{id}", handler.ViewId)
	mux.HandleFunc("GET /api/v1/categories", handler.ViewAll)

	return mux
}
