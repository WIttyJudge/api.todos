package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/usecase"
)

type TodoController interface {
	AllTodos() http.HandlerFunc
	CreateTodo() http.HandlerFunc
	DeleteTodo() http.HandlerFunc
}

type todoController struct {
	usecase usecase.TodoUsecase
}

func NewTodoController(usecase usecase.TodoUsecase) TodoController {
	return &todoController{usecase}
}

func (c *todoController) AllTodos() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-type", "application/json")

		// todos := entities.Todo{Title: "test", Task: "test"}
		todos, err := c.usecase.FetchAll()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(todos)
	}
}

func (c *todoController) CreateTodo() http.HandlerFunc {
	todo := &entities.Todo{}

	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")

		if err := json.NewDecoder(req.Body).Decode(todo); err != nil {
			resp.WriteHeader(http.StatusNoContent)
			resp.Write([]byte(`{"error": "Cannot parse data to json"}`))
			return
		}

		err := c.usecase.Store(todo)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"error": "Cannot create a new todo"}`))
			return
		}

		resp.WriteHeader(http.StatusCreated)
		resp.Write([]byte(`{"ok": "Good"}`))

	}
}

func (c *todoController) DeleteTodo() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-type", "application/json")

		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`Error`))
			return
		}

		deleted, err := c.usecase.Delete(id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Error"))
		}

		fmt.Print(deleted)
	}
}
