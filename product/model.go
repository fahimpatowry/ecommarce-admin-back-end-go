package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL        []string           `bson:"url" json:"url"`
	Name       string             `bson:"name" json:"name"`
	Decription string             `bson:"description" json:"description"`
	CategoryID float64            `bson:"categoryId" json:"categoryId"` // FK
	Price      float64            `bson:"price" json:"price"`
	OrderCount int                `bson:"orderCount" json:"orderCount"`
	Discount   float64            `bson:"discount" json:"discount"`
	Tag        string             `bson:"tag" json:"tag"`
	IsPopular  bool               `bson:"isPopular" json:"isPopular"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}
