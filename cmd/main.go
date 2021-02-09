package main

import (
	"log"

	"github.com/yongDataScince/rest-api"
	"github.com/yongDataScince/rest-api/pkg/handler"
)

func main() {
	srv := new(rest.Server)
	handlers := new(handler.Handlers)
	if err := srv.Run("5000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}