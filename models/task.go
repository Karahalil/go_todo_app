package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // e.g., "pending", "in-progress",
	UserID      int    `json:"-"`      // Foreign key to associate task with a user
}
