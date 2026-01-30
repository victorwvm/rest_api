package models

type Film struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate int     `json:"release_date"`
	Platform    string  `json:"platform"`
	Rating      float64 `json:"rating"`
}
