package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrizalr-remind-me/go-backend/internal/server"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New()
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
