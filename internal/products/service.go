package products

import (
	"context"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
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
