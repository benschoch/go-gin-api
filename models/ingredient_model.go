package models

type Ingredient struct {
	Id       string   `json:"id" bson:"_id"`
	Language string   `json:"language,omitempty" bson:"language,omitempty"`
	Singular string   `json:"singular,omitempty" bson:"singular,omitempty"`
	Plural   string   `json:"plural,omitempty" bson:"plural,omitempty"`
	Synonyms []string `json:"synonyms,omitempty" bson:"synonyms,omitempty"`
}
