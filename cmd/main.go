package main

import (
	"ToDo-List/pkg/handler"
	"ToDo-List/pkg/server"
	"ToDo-List/pkg/service"
	"ToDo-List/pkg/storage"
	"log"
)

func main() {
	rep := storage.NewPostgresStorage()
	services := service.NewServices(rep)
	api := handler.NewHandler(services)

	srv := server.NewServer()
	err := srv.Run(api.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
