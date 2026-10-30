package catalog

import "context"

type CatalogService interface {
	List(ctx context.Context) ([]CatalogWebResponse, error)
	Get(ctx context.Context, id int) (CatalogWebResponse, error)
	Create(ctx context.Context, catalog CatalogResponse) (CatalogWebResponse, error)
	Delete(ctx context.Context, categoryId int) error
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

func (s *CatalogServiceImpl) Get(ctx context.Context, id int) (CatalogWebResponse, error) {
	rawResponse, err := s.Repository.Get(ctx, id)

	if err != nil {
		return CatalogWebResponse{}, err
	}

	return ToWebResponse(rawResponse), nil
}

func (s *CatalogServiceImpl) Create(ctx context.Context, catalog CatalogResponse) (CatalogWebResponse, error) {
	rawResponse, err := s.Repository.Create(ctx, catalog)

	if err != nil {
		return CatalogWebResponse{}, err
	}

	return ToWebResponse(rawResponse), nil
}

func (s *CatalogServiceImpl) Delete(ctx context.Context, categoryId int) error {
	category, err := s.Repository.Get(ctx, categoryId)

	if err != nil {
		return err
	}
	s.Repository.Delete(ctx, category)
	return nil
}
