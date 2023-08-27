package routes

import (
	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	SignUp() gin.HandlerFunc
	SignIn() gin.HandlerFunc
	GetQuote() gin.HandlerFunc
	GetAllQuotes() gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}

type Handlers struct {
	authSvc  proto.AuthenticationServiceClient
	quoteSvc proto.QuoteServiceClient
}

func NewHandlers(authSvc proto.AuthenticationServiceClient, quoteSvc proto.QuoteServiceClient) HandlersInterface {
	return &Handlers{
		authSvc:  authSvc,
		quoteSvc: quoteSvc,
	}
}
