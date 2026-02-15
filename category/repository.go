package category

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
		collection: db.Collection("categorys"),
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]Category, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var category []Category

	if err := cursor.All(ctx, &category); err != nil {
		return nil, err
	}

	return category, nil
}

func (r *Repository) Create(ctx context.Context, c *Category) error {
	_, err := r.collection.InsertOne(ctx, c)

	return err
}

func (r *Repository) Update(ctx context.Context, id primitive.ObjectID, c *Category) error {
	update := bson.M{
		"$set": bson.M{
			"url":     c.URL,
			"slug":    c.Slug,
			"updated": c.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateByID(ctx, id, update)

	return err
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
