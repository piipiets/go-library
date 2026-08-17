package service

import (
	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/repository"
)

type CategoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(
	repository *repository.CategoryRepository,
) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) CreateCategory(
	categoryRequest request.CategoryRequest,
	username string,
) (model.Categories, error) {
	return s.repository.Create(categoryRequest, username)
}

func (s *CategoryService) GetAllCategories() ([]model.Categories, error) {
	return s.repository.GetAllCategory()
}
