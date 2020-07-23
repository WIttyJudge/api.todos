package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// User is an entity of users database table.
// Nickname and password is required.
type User struct {
	ID                int        `json:"id"`
	Nickname          string     `json:"nickname"`
	Password          string     `json:"password"`
	EncryptedPassword string     `json:"-" db:"encrypted_password"`
	CreatedAt         *time.Time `json:"created_at" db:"created_at"`
}

// Validate validates User struct fields.
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Nickname, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 40)),
	)
}
