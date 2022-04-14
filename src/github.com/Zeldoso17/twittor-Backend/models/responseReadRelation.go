package models

// ResponseReadRelation is a struct that contains true or false it is used to send a response to the client
type ResponseReadRelation struct {
	Status bool `json:"status"`
}