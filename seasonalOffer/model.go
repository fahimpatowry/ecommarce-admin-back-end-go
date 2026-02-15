package seasonalOffer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeasonalOffer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL       string             `bson:"url" json:"url"`
	Slug      string             `bson:"slug" json:"slug"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
	Position  int                `bson:"position" json:"position"`
	CreateAt  time.Time          `bson:"createAt" json:"createAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`
}
