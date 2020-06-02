package entities

import "time"

type Todos struct {
	ID int
	Title string
	Body string
	Completed bool
	CreatedAt time.Time
}