package main

import (
	"log"
	"todo"
	"todo/pkg/handlers"
)

func main() {
	handler := new(handlers.Handler)
	srv := new(todo.Server)
	err := srv.Run("8000", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error occurred while runnin http server: %s", err.Error())
	}
}
