package server

import (
	"github.com/aalsa16/golang-microservices/api/middleware"
	"github.com/aalsa16/golang-microservices/api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	handlers   routes.HandlersInterface
}

func NewServer(listenAddr string, handlers routes.HandlersInterface) *Server {
	return &Server{
		listenAddr: listenAddr,
		handlers:   handlers,
	}
}

func (s *Server) Run() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/signup", s.handlers.SignUp())
		auth.POST("/signin", s.handlers.SignIn())
		auth.POST("/refreshToken", s.handlers.RefreshToken())
	}

	quotes := router.Group("/quotes").Use(middleware.Auth())
	{
		quotes.GET("/getQuote", s.handlers.GetQuote())
		quotes.GET("/getAllQuotes", s.handlers.GetAllQuotes())
	}

	router.Run(s.listenAddr)
}
