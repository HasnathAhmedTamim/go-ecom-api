package models

type Order struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	Products  map[string]int `json:"products"`
	Total     float64        `json:"total"`
	Status    string         `json:"status"`
	CreatedAt string         `json:"created_at"`
}
