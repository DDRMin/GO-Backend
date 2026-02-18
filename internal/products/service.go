package products

import (
	"context"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, arg repo.CreateProductParams) (int64, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}

func (s *service) CreateProduct(ctx context.Context, arg repo.CreateProductParams) (int64, error) {
	id, err := s.repo.CreateProduct(ctx, arg)
	if err != nil {
		return 0, err
	}

	return id, nil
}
