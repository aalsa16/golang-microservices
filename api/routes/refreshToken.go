package routes

import (
	"context"
	"net/http"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

type RefreshData struct {
	Refresh_token string `json:"refresh_token"`
}

func (h *Handlers) RefreshToken() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var body RefreshData

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if body.Refresh_token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token cannot be blank"})
			c.Abort()
			return
		}

		resp, err := h.authSvc.RefreshToken(context.Background(), &proto.RefreshTokenRequest{
			RefreshToken: body.Refresh_token,
		})

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while refreshing token"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": resp})
	}

	return gin.HandlerFunc(fn)
}
