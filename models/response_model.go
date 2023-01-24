package models

type ApiResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ApiResponse2 struct {
	Status         int                    `json:"status"`
	Message        string                 `json:"message"`
	Total          int64                  `json:"total"`
	PageNumber     int64                  `json:"page_number"`
	RecipesPerPage int64                  `json:"recipes_per_page"`
	Data           map[string]interface{} `json:"data"`
}
