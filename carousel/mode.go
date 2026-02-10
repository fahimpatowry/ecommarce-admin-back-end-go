package carousel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Carousel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL string `bson:"url" json:"url"`
	Slug string `bson:"Slug" json:"Slug"`
	IsActive bool `bson:"isActive" json:'isActive'`
	CreateAt time.Time `bson:"createAt" json:"createAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}