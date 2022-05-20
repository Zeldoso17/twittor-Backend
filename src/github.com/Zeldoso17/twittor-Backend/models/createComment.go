package models

import (
	"time"
)

type CreateComment struct {
	UserID string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
	TweetID string `bson:"tweetid" json:"tweetid,omitempty"`
}