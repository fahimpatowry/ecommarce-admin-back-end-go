package product

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("products"),
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]Product, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var product []Product

	if err := cursor.All(ctx, &product); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) Create(ctx context.Context, c *Product) error {
	_, err := r.collection.InsertOne(ctx, c)

	return err
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, c *Product) error {
	update := bson.M{
		"$set": bson.M{
			"title":      c.Title,
			"decription": c.Decription,
			"tag":        c.Tag,
			"updated":    c.UpdatedAt,
			"url":        c.URL,
			"isPopular":  c.IsPopular,
			"updatedAt":  c.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateByID(ctx, id, update)

	return err
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
