package orders

import (
	"context"
	"errors"

	repo "github.com/Piyush-Singh-coder/ecom/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	PlaceOrder(ctx context.Context, temp createOrderParams) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {

	if tempOrder.CustomerId == 0 {
		return repo.Order{}, errors.New("customer id is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, errors.New("order items are required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)


	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerId)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.Items{
		product, err := qtx.FindProductById(ctx, item.ProductId)
		if err != nil {
			return repo.Order{}, errors.New("product not found")
		}

		if product.Quantity < int32(item.Quantity) {
			return repo.Order{}, errors.New("product out of stock")
		}

		_,err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID: order.ID,
			ProductID: item.ProductId,
			Quantity: int32(item.Quantity),
			PriceCents: product.PriceInCents,
		})

		if err != nil {
			return repo.Order{}, err
		}
	}
	
	if err := tx.Commit(ctx); err != nil {
		return repo.Order{}, err
	}

	return order, nil
}
