package product

import (
	"context"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
	productRepo "github.com/KeviinMoralees/crud-products/pkg/domains/repository/product"
)

type Service interface {
	CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	GetAllProducts(ctx context.Context) ([]entities.Product, error)
	GetProductByID(ctx context.Context, id int64) (entities.Product, error)
	UpdateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type service struct {
	repository productRepo.Repository
}

func NewService(repository productRepo.Repository) Service {
	return &service{repository: repository}
}
