package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RevanAdr/go-restful-api/internal/repository"
)

type Handler struct {
	// soon will be middleware here
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ViewId(w http.ResponseWriter, r *http.Request) {
	idstring := r.PathValue("id")

	id, _ := strconv.Atoi(idstring)

	w.Header().Set("Content-Type", "application/json")

	response := repository.GetById(id)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Fatal("Cannot encode data into response")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ViewAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	responses := repository.GetAll()

	err := json.NewEncoder(w).Encode(responses)

	if err != nil {
		log.Fatal("Cannot encode data to json responses")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
