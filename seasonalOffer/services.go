package seasonalOffer

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

func (s *Service) GetSeasonalOffers(ctx context.Context) ([]SeasonalOffer, error) {
	// return s.repo.GetAll(ctx)
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateSeasonalOffer(ctx context.Context, c *SeasonalOffer) error {
	if c.URL == "" || c.Slug == "" {
		return errors.New("url and slug are required")
	}

	now := time.Now()
	c.ID = primitive.NewObjectID()
	c.CreateAt = now
	c.UpdatedAt = now

	return s.repo.Create(ctx, c)
}

func (s *Service) UpdateSeasonalOffer(ctx context.Context, id primitive.ObjectID, c *SeasonalOffer) error {

	now := time.Now()
	c.UpdatedAt = now

	return s.repo.Update(ctx, id, c)
}

func (s *Service) DeleteSeasonalOffer(ctx context.Context, id primitive.ObjectID) error {

	return s.repo.Delete(ctx, id)
}
