package model

type Chapter struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Course   *Course   `json:"course"`
	Category *Category `json:"category"`
}

type NewChapter struct {
	Name     string `json:"name"`
	CourseID string `json:"courseId"`
}
