package models

type GiveLike struct {
	UserID string `bson:"userid" json:"userid,omitempty"`
    TweetID string `bson:"tweetid" json:"tweetid,omitempty"`
}