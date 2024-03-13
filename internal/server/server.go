package server

import (
	"camarinb2096/form_example/internal/app/form"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.TestMode)
	}

	router := gin.Default()
	s := &Server{router: router}
	s.routes(db)

	return s
}

func configCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (s *Server) routes(db *gorm.DB) {

	formSrv := form.NewService(db)

	endpoint := form.NewEndpoints(formSrv)

	s.router.Use(configCors())

	s.router.POST("/form", func(c *gin.Context) {
		endpoint.Post(c)
	})
}

func (s *Server) Start() {
	log.Println("Starting server on port 8080 on:", os.Getenv("GIN_MODE"))
	s.router.Run()
}
