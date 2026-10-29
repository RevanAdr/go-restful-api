package catalog

import "context"

type CatalogService interface {
	List(ctx context.Context) ([]CatalogWebResponse, error)
}

type CatalogServiceImpl struct {
	Repository CatalogRepository
}

func NewCatalogService(repo CatalogRepository) CatalogService {
	return &CatalogServiceImpl{Repository: repo}
}

func (s *CatalogServiceImpl) List(ctx context.Context) ([]CatalogWebResponse, error) {
	rawResponses, err := s.Repository.List(ctx)

	if err != nil {
		return []CatalogWebResponse{}, err
	}

	return ToWebResponses(rawResponses), nil
}
