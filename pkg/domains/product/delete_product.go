package product

import "context"

func (s *service) DeleteProduct(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
