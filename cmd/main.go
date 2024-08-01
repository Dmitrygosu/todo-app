package main

import (
	todo_app "github.com/Dmitrygosu/todo-app"
	"github.com/Dmitrygosu/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s ", err.Error())
	}
}
