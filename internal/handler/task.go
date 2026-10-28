package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RevanAdr/go-restful-api/internal/service"
)

type Handler interface {
	ViewId(w http.ResponseWriter, r *http.Request)
	ViewAll(w http.ResponseWriter, r *http.Request)
}

type HandlerImpl struct {
	service service.CategoryService
}

func NewHandler(service service.CategoryService) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (h *HandlerImpl) ViewId(w http.ResponseWriter, r *http.Request) {
	idstring := r.PathValue("id")

	id, _ := strconv.Atoi(idstring)

	w.Header().Set("Content-Type", "application/json")

	response := h.service.GetByIdService(id)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Fatal("Cannot encode data into response")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HandlerImpl) ViewAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	responses := h.service.GetAllService()

	err := json.NewEncoder(w).Encode(responses)

	if err != nil {
		log.Fatal("Cannot encode data to json responses")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
