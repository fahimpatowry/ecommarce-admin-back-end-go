package product

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetProducts(ctx context.Context) ([]Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateProduct(ctx context.Context, c *Product) error {
	// if c.URL == "" {
	// 	return errors.New("url and slug are required")
	// }

	now := time.Now()
	c.ID = primitive.NewObjectID()
	c.CreateAt = now
	c.UpdatedAt = now

	return s.repo.Create(ctx, c)
}

func (s *Service) UpdateProduct(ctx context.Context, id primitive.ObjectID, c *Product) error {
	now := time.Now()
	c.UpdatedAt = now

	return s.repo.Update(ctx, id, c)
}

func (s *Service) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {

	return s.repo.Delete(ctx, id)
}
