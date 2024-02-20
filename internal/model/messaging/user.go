package messaging

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the messaging system.
type User struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	GoogleId            string             `bson:"googleId"`
	UserName            string             `bson:"userName"`
	AuthorizationLevel  int                `bson:"authorizationLevel"`
}
