package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handlers"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("error occurred while runnin http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
