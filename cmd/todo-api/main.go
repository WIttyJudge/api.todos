package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/controller"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/store/connection"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/store/postgres"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/usecase"
)

var (
	postgresConn                             = connection.Connect()
	todoRepo       repository.TodoRepository = postgres.NewPostgresTodo(postgresConn)
	todoUsecase    usecase.TodoUsecase       = usecase.NewTodoUsecase(todoRepo)
	todoController controller.TodoController = controller.NewTodoController(todoUsecase)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todos", todoController.AllTodos()).Methods("GET")
	r.HandleFunc("/api/todos", todoController.CreateTodo()).Methods("POST")
	r.HandleFunc("/api/todos/{id}", todoController.DeleteTodo()).Methods("DELETE")

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
