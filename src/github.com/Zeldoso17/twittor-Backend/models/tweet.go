package models

// Tweet captures the message of the tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}