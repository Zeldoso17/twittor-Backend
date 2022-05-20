package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnFollowersTweets is a struct that allows to return the tweets of all my followers
type ReturnFollowersTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UserRelationID string             `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Tweet          struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
