package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Content   *string            `json:"content" bson:"content" validate:"required,min=2,max=350"`
	SenderID  *string            `json:"sender_id" bson:"sender_id" validate:"required"`
	RoomID    *string            `json:"room_id" bson:"room_id" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
