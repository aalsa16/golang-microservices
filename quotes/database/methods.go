package database

import (
	"github.com/aalsa16/golang-microservices/proto"
	"github.com/aalsa16/golang-microservices/quotes/types"
)

func (d *SqlServer) SaveQuote(req *types.SaveQuote) (savedQuote *types.Quote, err error) {
	query := "INSERT INTO `quotes` (`owner_uuid`, `quote`, `author`) VALUES (?, ?, ?)"

	result, err := d.DB.Exec(query, req.Owner_uuid, req.Quote, req.Author)

	if err != nil {
		return savedQuote, err
	}

	lastInsertID, _ := result.LastInsertId()

	var row types.Quote
	query = "SELECT id, quote, author, created_at, owner_uuid FROM quotes WHERE id = ?"
	err = d.DB.QueryRow(query, lastInsertID).Scan(&row.ID, &row.Quote, &row.Author, &row.CreatedAt, &row.Owner_uuid)

	if err != nil {
		return savedQuote, err
	}

	return &row, nil
}

func (d *SqlServer) GetAllUserQuotes(uuid string) (quotes []*proto.GetQuoteResponse, err error) {
	res, err := d.DB.Query("SELECT quote, author, created_at FROM quotes WHERE quotes.owner_uuid = ?", uuid)

	if err != nil {
		return quotes, err
	}

	defer res.Close()

	var rowSlice []*proto.GetQuoteResponse
	for res.Next() {
		var row proto.GetQuoteResponse
		err := res.Scan(&row.Quote, &row.Author, &row.CreatedAt)

		if err != nil {
			return quotes, err
		}

		rowSlice = append(rowSlice, &row)
	}

	return rowSlice, nil
}
