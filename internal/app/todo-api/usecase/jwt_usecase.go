package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/utils"
)

type JWTUsecase interface {
	GenerateJWT(user *entities.User) (string, error)
	ValidateJWT()
}

type jwtUsecase struct {
	SecretKey string
}

type JWTCustomeClaims struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`

	jwt.StandardClaims
}

func NewJWTUsecase() JWTUsecase {
	return &jwtUsecase{
		SecretKey: utils.GetEnv("JWT_KEY"),
	}
}

func (u *jwtUsecase) GenerateJWT(user *entities.User) (string, error) {
	claims := JWTCustomeClaims{
		ID:       user.ID,
		Nickname: user.Nickname,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "api.todo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *jwtUsecase) ValidateJWT() {

}
