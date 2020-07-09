package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/usecase"
)

// TodoController is a business logic of todos entity.
type TodoController interface {
	AllTodos() http.HandlerFunc

	CreateTodo() http.HandlerFunc
	DeleteTodo() http.HandlerFunc
}

type todoController struct {
	usecase usecase.TodoUsecase
}

// NewTodoController inits business logic
func NewTodoController(usecase usecase.TodoUsecase) TodoController {
	return &todoController{usecase}
}

func (c *todoController) AllTodos() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")

		todos, err := c.usecase.FetchAll()
		if err != nil {
			// resp.WriteHeader(http.StatusInternalServerError)
			c.error(resp, req, http.StatusInternalServerError, err)
			return
		}

		// resp.WriteHeader(http.StatusOK)
		// json.NewEncoder(resp).Encode(todos)
		c.response(resp, req, http.StatusOK, todos)
	}
}

func (c *todoController) CreateTodo() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
		Task  string `json:"task"`
	}
	r := &request{}

	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")

		if err := json.NewDecoder(req.Body).Decode(r); err != nil {
			// c.error(resp, req, http.StatusBadRequest, err)
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		todo := &entities.Todo{
			Title: r.Title,
			Task:  r.Task,
		}

		err := c.usecase.Store(todo)
		if err != nil {
			// c.error(resp, req, http.StatusInternalServerError, err)
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		c.response(resp, req, http.StatusCreated, nil)
	}
}

func (c *todoController) DeleteTodo() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")

		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			c.error(resp, req, http.StatusBadRequest, err)
			return
		}

		_, err = c.usecase.Delete(id)
		if err != nil {
			c.error(resp, req, http.StatusNoContent, err)
			return
		}

		resp.WriteHeader(http.StatusOK)
	}
}

func (c *todoController) error(resp http.ResponseWriter, req *http.Request, code int, err error) {
	c.response(resp, req, code, map[string]string{"error": err.Error()})
}

func (c *todoController) response(resp http.ResponseWriter, req *http.Request, code int, data interface{}) {
	resp.WriteHeader(code)
	if data != nil {
		json.NewEncoder(resp).Encode(data)
	}
}
