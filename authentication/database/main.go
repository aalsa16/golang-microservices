package database

import (
	"database/sql"

	"github.com/aalsa16/golang-microservices/authentication/types"
)

type SqlServer struct {
	DB *sql.DB
}

type Methods interface {
	SaveUser(req *types.SignUpRequest) (signupUser *types.SaveResponse, err error)
	GetUser(email string) (user types.User, err error)
}
