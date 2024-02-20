package messaging

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Message represents a message in the system, including dice rolls.
type Message struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	User          User               `bson:"user"`
	MessageBody   string             `bson:"messageBody"`
	DiceRoll      string             `bson:"diceRoll,omitempty"`
	DiceResult    string             `bson:"diceResult,omitempty"`
	MessageStatus bool               `bson:"messageStatus,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty"`
}
