package routes

import (
	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	SignUp() gin.HandlerFunc
}

type Handlers struct {
	svc proto.AuthenticationServiceClient
}

func NewHandlers(svc proto.AuthenticationServiceClient) HandlersInterface {
	return &Handlers{
		svc: svc,
	}
}
