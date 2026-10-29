package catalog

import (
	"encoding/json"
	"net/http"
)

type ErrorOrList interface {
	[]CatalogWebResponse | map[string]string
}

func WriteJson[T ErrorOrList](w http.ResponseWriter, status int, responses T) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
