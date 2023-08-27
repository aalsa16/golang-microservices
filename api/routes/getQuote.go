package routes

import (
	"context"
	"net/http"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetQuote() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		val, provided := c.GetQuery("uuid")
		if !provided || len(val) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "uuid value must be provided"})
			c.Abort()
			return
		}

		res, err := h.quoteSvc.GetQuote(context.Background(), &proto.GetQuoteRequest{
			Uuid: val,
		})

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while creating quote"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": res})
	}

	return gin.HandlerFunc(fn)
}
