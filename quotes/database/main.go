package database

import (
	"database/sql"

	"github.com/aalsa16/golang-microservices/quotes/types"
)

type SqlServer struct {
	DB *sql.DB
}

type Methods interface {
	SaveQuote(req *types.SaveQuote) (savedQuote *types.Quote, err error)
	GetAllUserQuotes(uuid string) (quotes []types.Quote, err error)
}
