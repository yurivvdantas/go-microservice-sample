package model

type Crypto_request struct {
	Name        string `json:"name" bson:"name"`
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}
