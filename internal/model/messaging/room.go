package messaging

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Room represents a chat room.
type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Topic     string             `bson:"topic"`
	Messages  []Message          `bson:"messages"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}
