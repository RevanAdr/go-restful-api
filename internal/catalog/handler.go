package catalog

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type CatalogHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type CatalogHandlerImpl struct {
	Service CatalogService
}

func NewCatalogHandler(service CatalogService) CatalogHandler {
	return &CatalogHandlerImpl{Service: service}
}

func (s *CatalogHandlerImpl) List(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	responses, err := s.Service.List(ctx)

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "could not list categories"})
		return
	}

	WriteJson(w, http.StatusOK, responses)
}

func (s *CatalogHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]string{"error": "unvalid id"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	result, err := s.Service.Get(ctx, id)

	switch {
	case errors.Is(err, ErrNotFound):
		WriteJson(w, http.StatusNotFound, map[string]string{"error": "category not found"})
	case err != nil:
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	default:
		WriteJson(w, http.StatusOK, result)
	}
}
