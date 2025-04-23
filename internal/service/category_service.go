package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"
	"training-project/pkg/util"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

type CategoryService interface {
	IsCategoryExisted(ctx context.Context, id int64) (bool, error)
	FindAllCategories(ctx context.Context) ([]model.Category, error)
	FindCategoryById(ctx context.Context, id int64) (*model.Category, error)
	CreateCategory(ctx context.Context, newCategory *dto.CategoryCreateReqDTO) error
	UpdateCategoryById(ctx context.Context, id int64, updatedCategory *dto.CategoryUpdateReqDTO) error
	DeleteCategoryById(ctx context.Context, id int64) error
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (categoryService *categoryService) IsCategoryExisted(ctx context.Context, id int64) (bool, error) {
	existed, err := categoryService.categoryRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (categoryService *categoryService) FindAllCategories(ctx context.Context) ([]model.Category, error) {
	return categoryService.categoryRepository.FindAll(ctx)
}

func (categoryService *categoryService) FindCategoryById(ctx context.Context, id int64) (*model.Category, error) {
	return categoryService.categoryRepository.FindById(ctx, id)
}

func (categoryService *categoryService) CreateCategory(ctx context.Context, categoryCreateReqDTO *dto.CategoryCreateReqDTO) error {
	newCategory := model.Category{
		Name:        categoryCreateReqDTO.Name,
		Description: categoryCreateReqDTO.Description,
	}
	return categoryService.categoryRepository.Create(ctx, &newCategory)
}

func (categoryService *categoryService) UpdateCategoryById(ctx context.Context, id int64, updatedCategory *dto.CategoryUpdateReqDTO) error {
	foundCategory, err := categoryService.FindCategoryById(ctx, id)
	if err != nil {
		return err
	}
	util.ApplyCategoryUpdate(foundCategory, updatedCategory)
	return categoryService.categoryRepository.UpdateById(ctx, id, foundCategory)
}

func (categoryService *categoryService) DeleteCategoryById(ctx context.Context, id int64) error {
	existed, err := categoryService.IsCategoryExisted(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of category is not valid")
	}
	return categoryService.categoryRepository.DeleteById(ctx, id)
}
