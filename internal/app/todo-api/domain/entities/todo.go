package entities

import "time"

// Todo in an entity of todos database table.
type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Task        string     `json:"task"`
	Completed   bool       `json:"completed"`
	CreatedAt   *time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
