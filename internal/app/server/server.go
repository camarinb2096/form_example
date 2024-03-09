package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {

	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.TestMode)
	}

	router := gin.Default()

	return &Server{router: router}
}

func (s *Server) Start() {
	log.Println("Starting server on port 8080 on:", os.Getenv("GIN_MODE"))
	s.router.Run()
}
