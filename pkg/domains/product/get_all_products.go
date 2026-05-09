package product

import (
	"context"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
)

func (s *service) GetAllProducts(ctx context.Context) ([]entities.Product, error) {
	return s.repository.GetAll(ctx)
}
