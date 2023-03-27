package main

import (
	"log"
	"todo"
	"todo/pkg/handlers"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run("8000", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error occurred while runnin http server: %s", err.Error())
	}
}
