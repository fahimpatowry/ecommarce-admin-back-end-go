package category

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	CategoryID float64            `bson:"categoryId" json:"categoryId"` // FK
	URL        string             `bson:"url" json:"url"`
	Slug       string             `bson:"slug" json:"slug"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}
