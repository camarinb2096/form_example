package main

import (
	"camarinb2096/form_example/internal/config/db"
	"camarinb2096/form_example/internal/config/server"
	"camarinb2096/form_example/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	logger := logger.NewLogger()

	err := godotenv.Load("../.env")
	if err != nil {
		logger.Error("Error loading .env file")
	}

	config := db.NewConfig()
	pgSqlConn := db.NewDb(config, logger)
	defer db.CloseDb(pgSqlConn, logger)

	server := server.NewServer(pgSqlConn)
	server.Start(logger)
}
