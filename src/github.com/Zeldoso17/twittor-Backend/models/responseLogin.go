package models

// ResponseLogin is a struct that contains the token that returns the user after login
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}