package products

import (
	"context"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) error
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) ListProducts(ctx context.Context) error {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return err
	}

	_ = products // TODO: Use products in handler

	return nil
}
