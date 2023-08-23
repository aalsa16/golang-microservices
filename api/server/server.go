package server

import (
	"github.com/aalsa16/golang-microservices/api/routes"
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

	auth := router.Group("/auth")
	{
		auth.POST("/signup", s.handlers.SignUp())
	}

	router.Run(s.listenAddr)
}
