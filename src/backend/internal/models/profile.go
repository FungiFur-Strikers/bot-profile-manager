// internal/models/profile.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BotID       string             `bson:"botId" json:"botId"`
	Name        string             `bson:"name" json:"name"`
	Personality string             `bson:"personality" json:"personality"`
	Avatar      string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Language    string             `bson:"language,omitempty" json:"language,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
