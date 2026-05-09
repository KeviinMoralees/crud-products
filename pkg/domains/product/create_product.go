package product

import (
	"context"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
)

func (s *service) CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	return s.repository.Create(ctx, product)
}
