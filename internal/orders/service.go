package orders

import (
	"context"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	CreateOrder(ctx context.Context, arg repo.CreateOrderParams) error
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrder(ctx context.Context, arg repo.CreateOrderParams) error {
	err := s.repo.CreateOrder(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}
