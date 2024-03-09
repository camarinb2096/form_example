package main

import (
	"camarinb2096/form_example/internal/app/config/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Failed to load .env file: ", err)
	}
	config := db.NewConfig()
	_ = db.NewDb(config)

	r := gin.Default()

	log.Println("Gin mode: ", os.Getenv("GIN_MODE"))

	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.TestMode)
	}

	r.POST("/pqrs", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
