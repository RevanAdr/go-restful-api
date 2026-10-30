package catalog

import (
	"encoding/json"
	"net/http"
)

type CatalogResponse struct {
	Id   int
	Name string
}

type CatalogWebResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CatalogCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

func ToWebResponses(raw []CatalogResponse) []CatalogWebResponse {
	results := []CatalogWebResponse{}
	for _, v := range raw {
		result := CatalogWebResponse{Id: v.Id, Name: v.Name}
		results = append(results, result)
	}
	return results
}

func ToWebResponse(raw CatalogResponse) CatalogWebResponse {
	return CatalogWebResponse{
		Id:   raw.Id,
		Name: raw.Name,
	}
}

func ReadBody(r *http.Request, result interface{}) {
	request := json.NewDecoder(r.Body)
	err := request.Decode(result)
	if err != nil {
		panic("error on reading request")
	}
}
