package orders

import (
	"context"
	"fmt"

	repo "github.com/DDRMin/GO-Backend/internal/adapters/sqlc"
)

type Service interface {
	CreateOrder(ctx context.Context, req CreateOrderRequest) (int64, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{repo: repo}
}

func (s *service) CreateOrder(ctx context.Context, req CreateOrderRequest) (int64, error) {
	if len(req.Items) == 0 {
		return 0, fmt.Errorf("order must have at least one item")
	}

	orderID, err := s.repo.CreateOrder(ctx, req.UserID)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	for _, item := range req.Items {
		err := s.repo.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
		if err != nil {
			return 0, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	return orderID, nil
}
