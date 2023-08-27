package routes

import (
	"context"
	"net/http"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/gin-gonic/gin"
)

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handlers) SignIn() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var body SignInData

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password cannot be empty fields"})
			c.Abort()
			return
		}

		user := &proto.SignInRequest{
			Email:    body.Email,
			Password: body.Password,
		}

		resp, err := h.authSvc.SignIn(context.Background(), user)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error while signing in"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": resp})
	}

	return gin.HandlerFunc(fn)
}
