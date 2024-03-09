package main

import (
	"camarinb2096/form_example/internal/app/config/db"
	"camarinb2096/form_example/internal/app/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Failed to load .env file: ", err)
	}

	config := db.NewConfig()

	pgSqlConn := db.NewDb(config)

	defer db.CloseDb(pgSqlConn)

	server := server.NewServer()
	server.Start()
}
