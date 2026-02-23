package models

type Order struct {
	ID       string   `json:"id"`
	UserID   string   `json:"user_id"`
	Products []string `json:"products"`
	Status   string   `json:"status"`
}
