package models

type Image struct {
	URL            string           `json:"url,omitempty" bson:"url,omitempty"`
	ImageVariation []ImageVariation `json:"variations,omitempty" bson:"variations,omitempty"`
}

type ImageVariation struct {
	Url    string `json:"url,omitempty" bson:"url,omitempty"`
	Filter string `json:"filter,omitempty" bson:"filter,omitempty"`
}
