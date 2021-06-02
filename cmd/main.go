package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/vladStepanenko1/todo-app"
	"github.com/vladStepanenko1/todo-app/pkg/handler"
	"github.com/vladStepanenko1/todo-app/pkg/repository"
	"github.com/vladStepanenko1/todo-app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     "192.168.43.135",
		Port:     "5432",
		Username: "postgres",
		Password: "keso",
		DBname:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("pisos: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
