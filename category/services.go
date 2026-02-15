package category

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetCategorys(ctx context.Context) ([]Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateCategory(ctx context.Context, c *Category) error {
	if c.URL == "" || c.Slug == "" {
		return errors.New("url and slug are required")
	}

	now := time.Now()
	c.ID = primitive.NewObjectID()
	c.CreateAt = now
	c.UpdatedAt = now

	return s.repo.Create(ctx, c)
}

func (s *Service) UpdateCategory(ctx context.Context, id primitive.ObjectID, c *Category) error {
	now := time.Now()
	c.UpdatedAt = now

	return s.repo.Update(ctx, id, c)
}

func (s *Service) DeleteCategory(ctx context.Context, id primitive.ObjectID) error {

	return s.repo.Delete(ctx, id)
}
