package service

import (
	"database/sql"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/aalsa16/golang-microservices/quotes/database"
)

type quoteService struct {
	database *database.SqlServer
	proto.UnimplementedQuoteServiceServer
}

func NewQuoteService(sql *sql.DB) proto.QuoteServiceServer {
	return &quoteService{
		database: &database.SqlServer{
			DB: sql,
		},
	}
}
