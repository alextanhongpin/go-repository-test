package repository

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/domain"
)

type ProductRepository struct {
	products *tables.ProductTable
}

func NewProductRepository(db atomic) *ProductRepository {
	return &ProductRepository{
		products: tables.NewProduct(db),
	}
}

func (r *ProductRepository) Find(ctx context.Context, id int64) (*domain.ProductOwner, error) {
	res, err := r.products.Find(ctx, id, tables.Relation{
		Name: "User",
	})
	if err != nil {
		return nil, fmt.Errorf("products.Find: %w", err)
	}

	return NewProductOwner(res), nil
}

func (r *ProductRepository) List(ctx context.Context) ([]*domain.Product, error) {
	res, err := r.products.List(ctx, &tables.OffsetPagination{
		Limit:   10,
		Offset:  0,
		OrderBy: []string{"id DESC"},
	})
	if err != nil {
		return nil, fmt.Errorf("products.List: %w", err)
	}

	toProduct := func(i int) *domain.Product {
		return NewProduct(&res[i])
	}

	return Map(res, toProduct), nil
}
