package catalog

import (
	"encoding/json"
	"net/http"
)

type ErrorOrList interface {
	[]CatalogWebResponse | map[string]string | CatalogWebResponse
}

func WriteJson[T ErrorOrList](w http.ResponseWriter, status int, responses T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responses)
}
