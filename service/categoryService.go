package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/repository"
)

var ErrCategoryNotFound = errors.New("category not found")

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

func (s *CategoryService) CreateCategory(categoryRequest request.CategoryRequest, username string) (response.CategoryResponse, error) {
	category, err := s.repository.Create(categoryRequest, username)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	return toCategoryResponse(category), nil
}

func (s *CategoryService) GetAllCategories() ([]response.CategoryResponse, error) {
	categories, err := s.repository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	categoryResponses := make([]response.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponses = append(categoryResponses, toCategoryResponse(category))
	}

	return categoryResponses, nil
}

func (s *CategoryService) GetCategoryById(id int) (response.CategoryResponse, error) {
	category, err := s.repository.GetCategoryById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.CategoryResponse{}, ErrCategoryNotFound
		}

		return response.CategoryResponse{}, err
	}

	return toCategoryResponse(category), nil
}

func toCategoryResponse(category model.Categories) response.CategoryResponse {
	return response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func (s *CategoryService) DeleteCategory(id int) error {
	err := s.repository.DeleteCategory(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCategoryNotFound
		}

		return err
	}

	return nil
}

func (s *CategoryService) GetBooksByCategory(id int) ([]response.BookResponse, error) {
	books, err := s.repository.GetBooksByCategory(id)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *CategoryService) UpdateCategory(id int, req request.CategoryRequest, username string) error {

	if req.Name == "" {
		return errors.New("name should be filled")
	}

	var category model.Categories
	category.Name = req.Name
	category.ModifiedAt = time.Now()
	category.ModifiedBy = username

	err := s.repository.UpdateCategory(id, category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCategoryNotFound
		}

		return err
	}

	return nil
}
