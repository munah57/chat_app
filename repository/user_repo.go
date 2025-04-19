package repository

import (
	"context"
	"real-chat/models"
	"time"

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

func (r *UserRepo) GetPaginatedUsers(ctx context.Context, recordsPerPage, page int) ([]models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

}
