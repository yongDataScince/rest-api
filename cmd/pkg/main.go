package main

import (
	"log"

	rest "github.com/yongDataScince/rest-api/cmd"
)

func main() {
	srv := new(rest.Server)
	if err := srv.Run("5000"); err != nil {
		log.Fatal(err)
	}

}