package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID                   primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID               primitive.ObjectID  `json:"userId" bson:"userId,omitempty"`
	RefreshToken         string              `json:"refreshToken"  bcon:"refreshToken"`
	ExpiresAt            time.Time           `json:"expiresAt"   bcon:"expiresAt"`
}