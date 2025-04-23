package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type categoryRepository struct {
	db *bun.DB
}

type CategoryRepository interface {
	ExistsById(ctx context.Context, id int64) (bool, error)
	FindAll(ctx context.Context) ([]model.Category, error)
	FindById(ctx context.Context, id int64) (*model.Category, error)
	Create(ctx context.Context, newCategory *model.Category) error
	UpdateById(ctx context.Context, id int64, updatedCategory *model.Category) error
	DeleteById(ctx context.Context, id int64) error
}

func NewCategoryRepository(db *bun.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (categoryRepository *categoryRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := categoryRepository.db.NewSelect().Model(&model.Category{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (categoryRepository *categoryRepository) FindAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := categoryRepository.db.NewSelect().Model(&categories).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (categoryRepository *categoryRepository) FindById(ctx context.Context, id int64) (*model.Category, error) {
	var category model.Category
	err := categoryRepository.db.NewSelect().Model(&category).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (categoryRepository *categoryRepository) Create(ctx context.Context, newCategory *model.Category) error {
	_, err := categoryRepository.db.NewInsert().Model(newCategory).Exec(ctx)
	return err
}

func (categoryRepository *categoryRepository) UpdateById(ctx context.Context, id int64, updatedCategory *model.Category) error {
	_, err := categoryRepository.db.NewUpdate().Model(updatedCategory).Where("id = ?", id).Exec(ctx)
	return err
}

func (categoryRepository *categoryRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := categoryRepository.db.NewDelete().Model(&model.Category{}).Where("id = ?", id).Exec(ctx)
	return err
}
