package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/aalsa16/golang-microservices/quotes/types"
)

func (s *quoteService) GetQuote(ctx context.Context, req *proto.GetQuoteRequest) (*proto.GetQuoteResponse, error) {
	resp, err := http.Get("https://api.quotable.io/random")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var quote types.QuoteResponse
	err = json.Unmarshal(body, &quote)

	if err != nil {
		return nil, err
	}

	savedQuote, err := s.database.SaveQuote(&types.SaveQuote{
		Owner_uuid: req.Uuid,
		Quote:      quote.Content,
		Author:     quote.Author,
	})

	if err != nil {
		return nil, err
	}

	return &proto.GetQuoteResponse{
		Quote:     savedQuote.Quote,
		Author:    savedQuote.Author,
		CreatedAt: savedQuote.CreatedAt,
	}, nil
}

func (s *quoteService) GetAllQuotes(req *proto.GetQuoteRequest, srv proto.QuoteService_GetAllQuotesServer) error {
	quotes, err := s.database.GetAllUserQuotes(req.Uuid)

	if err != nil {
		return err
	}

	for _, quote := range quotes {
		err = srv.Send(quote)

		if err != nil {
			return err
		}
	}

	return nil
}
