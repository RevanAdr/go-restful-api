package catalog

type CatalogResponse struct {
	Id   int
	Name string
}

type CatalogWebResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToWebResponses(raw []CatalogResponse) []CatalogWebResponse {
	results := []CatalogWebResponse{}
	for _, v := range raw {
		result := CatalogWebResponse{Id: v.Id, Name: v.Name}
		results = append(results, result)
	}
	return results
}
