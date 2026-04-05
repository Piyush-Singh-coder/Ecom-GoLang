package products

import (
	"context"

	repo "github.com/Piyush-Singh-coder/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	ProductById(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *svc) ProductById(ctx context.Context, id int64) (repo.Product, error){
	product, err := s.repo.FindProductById(ctx, id)

	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}
