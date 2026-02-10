package carousel

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service{
	return &Service{repo: r}
}

func (s *Service) GetCarousels(ctx context.Context) ([]Carousel, error){
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateCarousel(ctx context.Context, c *Carousel) error {
	if c.URL == "" || c.Slug == "" {
		return errors.New("url and slug are required")
	}

	now := time.Now()
	c.ID = primitive.NewObjectID()
	c.CreateAt = now
	c.UpdatedAt = now

	return s.repo.Create(ctx, c)
}

func (s *Service) UpdateCarousel(ctx context.Context, id primitive.ObjectID, c *Carousel) error {

	// if c.ID == nil {
	// 	return errors.New("url and slug are required")
	// }

	now := time.Now()
	c.UpdatedAt = now

	return s.repo.Update(ctx, id, c)
}

func (s *Service) DeleteCarousel(ctx context.Context, id primitive.ObjectID,) error {

	return s.repo.Delete(ctx, id)
}

