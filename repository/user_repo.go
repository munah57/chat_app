package repository

import (
	"context"
	"real-chat/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

type UserRepository interface {
	GetPaginatedUsers(ctx context.Context, recordsPerPage, page int) ([]models.User, error)
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepo{collection: collection}
}

type PaginatedUserResult struct {
	TotalCount int          `bson:"total_count"`
	Users      []models.User `bson:"user_items"`
}

func (r *UserRepo) GetPaginatedUsers(parentCtx context.Context, recordsPerPage, page int) ([]models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	startIndex := (page - 1) * recordsPerPage

	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	projectStage := bson.D{
		{Key: "_id", Value: 0},
		{Key: "total_count", Value: 1},
		{Key: "user_items", Value: bson.D{{Key: "$slice", Value: bson.A{"$data", startIndex, recordsPerPage}}}},
	}

	result, err := r.collection.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	if err != nil {
		return nil, err
	}

	var results []PaginatedUserResult
	if err = result.All(ctx, &results); err != nil {
		return nil, err
	}

	return results[0].Users, nil

}
