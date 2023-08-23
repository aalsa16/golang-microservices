package service

import (
	"database/sql"

	"github.com/aalsa16/golang-microservices/authentication/database"
	"github.com/aalsa16/golang-microservices/proto"
)

type authService struct {
	database *database.SqlServer
	proto.UnimplementedAuthenticationServiceServer
}

func NewAuthService(sql *sql.DB) proto.AuthenticationServiceServer {
	return &authService{
		database: &database.SqlServer{
			DB: sql,
		},
	}
}
