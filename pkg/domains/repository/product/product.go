package product

import (
	"context"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
)

type Repository interface {
	Create(ctx context.Context, product entities.Product) (entities.Product, error)
	GetAll(ctx context.Context) ([]entities.Product, error)
	GetByID(ctx context.Context, id int64) (entities.Product, error)
	Update(ctx context.Context, product entities.Product) (entities.Product, error)
	Delete(ctx context.Context, id int64) error
}
