package main

import (
	"camarinb2096/form_example/internal/server"
	"camarinb2096/form_example/pkg/infra/db"
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

	server := server.NewServer(pgSqlConn)
	server.Start()
}
