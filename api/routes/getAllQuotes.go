package routes

import (
	"context"
	"io"
	"net/http"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

type Quotes struct {
	Quote     string `json:"quote"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}

func (h *Handlers) GetAllQuotes() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		val, provided := c.GetQuery("uuid")
		if !provided || len(val) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "uuid value must be provided"})
			c.Abort()
			return
		}

		var quotes []Quotes

		res, err := h.quoteSvc.GetAllQuotes(context.Background(), &proto.GetQuoteRequest{
			Uuid: val,
		})

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while getting all quotes"})
			c.Abort()
			return
		}

		for {
			quoteres, err := res.Recv()

			if err == io.EOF {
				err = res.CloseSend()

				if err != nil {
					c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while getting all quotes"})
					c.Abort()
					return
				}

				break
			}

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while getting all quotes"})
				c.Abort()
				return
			}

			quotes = append(quotes, Quotes{
				Quote:     quoteres.Quote,
				Author:    quoteres.Author,
				CreatedAt: quoteres.CreatedAt,
			})
		}

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while getting all quotes"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": quotes})
	}

	return gin.HandlerFunc(fn)
}
