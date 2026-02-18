package orders

import (
	"context"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	CreateOrder(ctx context.Context, arg repo.Querier) (int64, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrder(ctx context.Context, arg repo.Querier) (int64, error) {
	orderID, err := s.repo.CreateOrder(ctx, 1) // Assuming userID is 1 for demonstration purposes
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

