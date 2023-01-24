package models

type Collection struct {
	ID       string `json:"id" bson:"_id"`
	Language string `json:"language,omitempty" bson:"language,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
}
