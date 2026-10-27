package handler

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewHandler()

	mux.HandleFunc("GET /api/v1/categories/{id}", handler.ViewId)
	mux.HandleFunc("GET /api/v1/categories", handler.ViewAll)

	return mux
}
