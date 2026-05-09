package product

import (
	"context"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
)

func (s *service) GetProductByID(ctx context.Context, id int64) (entities.Product, error) {
	return s.repository.GetByID(ctx, id)
}
