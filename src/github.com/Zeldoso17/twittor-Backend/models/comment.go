package models

type Comment struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}