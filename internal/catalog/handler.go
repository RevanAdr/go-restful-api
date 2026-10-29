package catalog

import "net/http"

type CatalogHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}

type CatalogHandlerImpl struct {
	Service CatalogService
}

func NewCatalogHandler(service CatalogService) CatalogHandler {
	return &CatalogHandlerImpl{Service: service}
}

func (s *CatalogHandlerImpl) List(w http.ResponseWriter, r *http.Request) {
	responses, err := s.Service.List(r.Context())

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "could not list products"})
		return
	}

	WriteJson(w, http.StatusOK, responses)
}
