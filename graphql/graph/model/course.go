package model

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Category    *Category `json:"category"`
}

type NewCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"categoryId"`
}
