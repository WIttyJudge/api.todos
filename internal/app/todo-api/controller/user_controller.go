package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/usecase"
)

type UserController interface {
	Login() http.HandlerFunc
	Signup() http.HandlerFunc
}

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return &userController{usecase}
}

func (c *userController) Login() http.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		r := &request{}

		resp.Header().Set("Content-type", "application/json")
		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		user := &entities.User{
			Nickname: r.Nickname,
			Password: r.Password,
		}

		user, err := c.usecase.Login(user)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusUnauthorized)
		}
	}
}

func (c *userController) Signup() http.HandlerFunc {
	user := &entities.User{}

	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")

		if err := json.NewDecoder(req.Body).Decode(user); err != nil {
			c.error(resp, http.StatusBadRequest, err)
			return
		}

		if err := c.usecase.Store(user); err != nil {
			c.error(resp, http.StatusUnprocessableEntity, err)
			return
		}

		c.response(resp, http.StatusCreated, nil)
	}
}

func (c *userController) error(resp http.ResponseWriter, code int, err error) {
	c.response(resp, code, map[string]string{"error": err.Error()})
}

func (c *userController) response(resp http.ResponseWriter, code int, data interface{}) {
	resp.WriteHeader(code)
	if data != nil {
		json.NewEncoder(resp).Encode(data)
	}
}
