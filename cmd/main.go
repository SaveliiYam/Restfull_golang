package main

import (
	"log"

	todo "example.com/m/v2"
	"example.com/m/v2/pkg/handler"
	"example.com/m/v2/pkg/repository"
	"example.com/m/v2/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
