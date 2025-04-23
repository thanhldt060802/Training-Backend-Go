package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type productRepository struct {
	db *bun.DB
}

type ProductRepository interface {
	ExistsById(ctx context.Context, id int64) (bool, error)
	FindAll(ctx context.Context) ([]model.Product, error)
	FindById(ctx context.Context, id int64) (*model.Product, error)
	Create(ctx context.Context, newProduct *model.Product) error
	UpdateById(ctx context.Context, id int64, updatedProduct *model.Product) error
	DeleteById(ctx context.Context, id int64) error
}

func NewProductRepository(db *bun.DB) ProductRepository {
	return &productRepository{db: db}
}

func (productRepository *productRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := productRepository.db.NewSelect().Model(&model.Product{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (productRepository *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := productRepository.db.NewSelect().Model(&products).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *productRepository) FindById(ctx context.Context, id int64) (*model.Product, error) {
	var product model.Product
	err := productRepository.db.NewSelect().Model(&product).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepository *productRepository) Create(ctx context.Context, newProduct *model.Product) error {
	_, err := productRepository.db.NewInsert().Model(newProduct).Exec(ctx)
	return err
}

func (productRepository *productRepository) UpdateById(ctx context.Context, id int64, updatedProduct *model.Product) error {
	_, err := productRepository.db.NewUpdate().Model(updatedProduct).Where("id = ?", id).Exec(ctx)
	return err
}

func (productRepository *productRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := productRepository.db.NewDelete().Model(&model.Product{}).Where("id = ?", id).Exec(ctx)
	return err
}
