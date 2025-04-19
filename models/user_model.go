package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" validate:"required,min=2,max=100"`
	FirstName    *string            `json:"first_name" bson:"first_name"`
	LastName     *string            `json:"last_name" bson:"last_name"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=6,max=64"`
	Email        *string            `json:"email" bson:"email" validate:"required,email"`
	Avatar       *string            `json:"avatar" bson:"avatar,omitempty"`
	Phone        *string            `json:"phone" bson:"phone" validate:"required"`
	Token        *string            `json:"token,omitempty" bson:"token,omitempty"`
	RefreshToken *string            `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	UserID       string             `json:"user_id" bson:"user_id"`
}
