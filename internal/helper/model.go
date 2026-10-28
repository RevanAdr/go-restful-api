package helper

import (
	"github.com/RevanAdr/go-restful-api/internal/domain"
)

func ToResponse(RawResponse CategoryResponse) domain.WebResponse {
	result := domain.WebResponse{
		Id:   RawResponse.Id,
		Name: RawResponse.Name,
	}

	return result
}

func ToResponses(RawResponses []CategoryResponse) []domain.WebResponse {
	var results []domain.WebResponse
	for _, e := range RawResponses {
		v := domain.WebResponse{
			Id:   e.Id,
			Name: e.Name,
		}
		results = append(results, v)
	}

	return results
}
