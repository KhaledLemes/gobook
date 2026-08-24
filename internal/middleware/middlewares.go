package middleware

import (
	"gobook/internal/auth"
	"gobook/internal/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("\n%s, %s, %s at %s", c.Request.Method, c.Request.URL, c.Request.Host, time.Now())
		next(c)
	}
}

func Autentica(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.ValidadeToken(c.Request); err != nil {
			responses.Err(c, http.StatusUnauthorized, err)
			return
		}
		next(c)
	}
}
