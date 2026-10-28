package service

import (
	"github.com/RevanAdr/go-restful-api/internal/domain"
	"github.com/RevanAdr/go-restful-api/internal/helper"
	"github.com/RevanAdr/go-restful-api/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService interface {
	GetByIdService(id int) domain.WebResponse
	GetAllService() []domain.WebResponse
}

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	db         *pgxpool.Pool
}

func NewCategoryService(repository repository.CategoryRepository, db *pgxpool.Pool) CategoryService {
	return &CategoryServiceImpl{
		repository: repository,
		db:         db,
	}
}

func (s *CategoryServiceImpl) GetByIdService(id int) domain.WebResponse {
	response := s.repository.GetById(id)

	return helper.ToResponse(*response)
}

func (s *CategoryServiceImpl) GetAllService() []domain.WebResponse {
	response := s.repository.GetAll()

	return helper.ToResponses(*response)
}
