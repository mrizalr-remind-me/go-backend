package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrizalr-remind-me/go-backend/internal/server"
	"github.com/mrizalr-remind-me/go-backend/pkg/db/mysql/sqlx"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env file - ", err)
	}

	db, err := sqlx.New()
	if err != nil {
		log.Fatal("error connect to database - ", err)
	}

	s := server.New(db)
	err = s.Run()
	if err != nil {
		log.Fatal("error running the server - ", err)
	}
}
