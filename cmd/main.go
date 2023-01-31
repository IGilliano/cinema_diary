package main

import (
	"cinema_diary"
	"cinema_diary/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(cinema_diary.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
