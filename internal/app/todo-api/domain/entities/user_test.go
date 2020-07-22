package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GenerateUser creates test valid user.
func GenerateUser(t *testing.T) *User {
	t.Helper()

	return &User{
		ID:       1,
		Nickname: "Validuser",
		Password: "ValidPassword",
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    func() *User
		wantErr bool
	}{
		{
			name: "valid",
			user: func() *User {
				return GenerateUser(t)
			},
			wantErr: false,
		},
		{
			name: "required nickname",
			user: func() *User {
				u := GenerateUser(t)
				u.Nickname = ""
				return u
			},
			wantErr: true,
		},
		{
			name: "required password",
			user: func() *User {
				u := GenerateUser(t)
				u.Password = ""
				return u
			},
			wantErr: true,
		},
		{
			name: "password length less then 6",
			user: func() *User {
				u := GenerateUser(t)
				u.Password = "test"
				return u
			},
			wantErr: true,
		},
		{
			name: "password length more then 40",
			user: func() *User {
				u := GenerateUser(t)
				// Here the length in 44
				u.Password = "testtesttesttesttesttesttesttesttesttesttest"
				return u
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			if tt.wantErr {
				assert.NotNil(tt.user().Validate())
			}

			if !tt.wantErr {
				assert.Nil(tt.user().Validate())
			}
		})
	}
}
