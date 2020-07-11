package entities

import "time"

// User is an entity of users database table
type User struct {
	ID                int        `json:"id"`
	Nickname          string     `json:"nickname"`
	Password          string     `json:"password"`
	EncryptedPassword string     `json:"-"`
	CreateAt          *time.Time `json:"created_at"`
}
