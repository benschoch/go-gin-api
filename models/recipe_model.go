package models

type Recipe struct {
	Id                    string                  `json:"id" bson:"_id"`
	Language              string                  `json:"language,omitempty" bson:"language,omitempty"`
	IsPublished           bool                    `json:"is_published,omitempty" bson:"is_published,omitempty"`
	Title                 string                  `json:"title,omitempty" bson:"title,omitempty"`
	Slug                  string                  `json:"slug,omitempty" bson:"slug,omitempty"`
	PreparationTime       int                     `json:"preparation_time,omitempty" bson:"preparation_time,omitempty"`
	CookingTime           int                     `json:"cooking_time,omitempty" bson:"cooking_time,omitempty"`
	Difficulty            int                     `json:"difficulty,omitempty" bson:"difficulty,omitempty"`
	YoutubeVideoId        string                  `json:"youtube_video_id,omitempty" bson:"youtube_video_id,omitempty"`
	Course                []Course                `json:"courses,omitempty" bson:"courses,omitempty"`
	Diet                  []Diet                  `json:"diets,omitempty" bson:"diets,omitempty"`
	RecipeIngredientGroup []RecipeIngredientGroup `json:"ingredient_groups,omitempty" bson:"ingredient_groups,omitempty"`
	ServingType           ServingType             `json:"serving_type,omitempty" bson:"serving_type,omitempty"`
}

type RecipeIngredientGroup struct {
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	Order            int                `json:"order,omitempty" bson:"order,omitempty"`
	RecipeIngredient []RecipeIngredient `json:"ingredients,omitempty" bson:"recipe_ingredients,omitempty"`
}

type RecipeIngredient struct {
	Ingredient       Ingredient `json:"ingredient,omitempty" bson:"ingredient,omitempty"`
	Unit             Unit       `json:"unit,omitempty" bson:"unit,omitempty"`
	AmountFrom       int        `json:"amount_from,omitempty" bson:"amount_from,omitempty"`
	AmountTo         int        `json:"amount_to,omitempty" bson:"amount_to,omitempty"`
	IsMainIngredient bool       `json:"is_main_ingredient,omitempty" bson:"is_main_ingredient,omitempty"`
}
