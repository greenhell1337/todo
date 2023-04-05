package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo"
	"todo/pkg/handlers"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error any loading variables: %s", err.Error())
	}
	db, err := repository.NewMyDB(repository.Config{
		Username: "root",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "todo",
	})
	if err != nil {
		logrus.Fatalf("not connected, error: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := &todo.Server{}
	err = srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("error occurred while runnin http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
