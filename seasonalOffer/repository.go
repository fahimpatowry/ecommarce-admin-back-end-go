package seasonalOffer

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("seasonalOffers"),
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]SeasonalOffer, error) {

	cursor, err := r.collection.Find(ctx, bson.M{"isActive": true})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	fmt.Println("cursor", cursor)

	var seasonalOffer []SeasonalOffer

	if err := cursor.All(ctx, &seasonalOffer); err != nil {
		return nil, err
	}

	return seasonalOffer, nil
}

func (r *Repository) Create(ctx context.Context, c *SeasonalOffer) error {
	_, err := r.collection.InsertOne(ctx, c)

	return err
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, c *SeasonalOffer) error {
	update := bson.M{
		"$set": bson.M{
			"url":      c.URL,
			"slug":     c.Slug,
			"isActive": c.IsActive,
			"updated":  c.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateByID(ctx, id, update)

	return err
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
