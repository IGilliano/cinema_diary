package main

import (
	"cinema_diary"
	"cinema_diary/pkg/handler"
	"cinema_diary/pkg/repository"
	service2 "cinema_diary/pkg/service"
	"log"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "456123789",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Cant initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	service := service2.NewService(rep)
	handlers := handler.NewHandler(service)
	srv := new(cinema_diary.Server)

	if err = srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
