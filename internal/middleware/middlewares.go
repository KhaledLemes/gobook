package middleware

import (
	"gobook/internal/auth"
	"gobook/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Printf("\n%s, %s, %s at %s", c.Request.Method, c.Request.URL, c.Request.Host, time.Now())
		c.Next()
	}
}

// Autentica serve somente para rotas, não deve ser usado para servir HTML
func Autentica() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.ValidadeToken(c); err != nil {
			responses.Err(c, http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}

// AutenticaProprietario serve para autenticar páginas de front-end que somente proprietários e admin possuem acesso
func AutenticaProprietario() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.ValidadeToken(c); err != nil {
			responses.ErrHTML(c, http.StatusBadRequest, "unauthorized.html", err)
			return
		}
		role, err := auth.PegarRoleUsuario(c)
		if err != nil {
			responses.ErrHTML(c, http.StatusBadRequest, "unauthorized.html", err)
			return
		}
		if role != "owner" && role != "admin" {
			responses.ErrHTML(c, http.StatusForbidden, "unauthorized.html", ErrRoleInvalido)
			return
		}
		c.Next()
	}
}

// AutenticaAdmin serve para autenticar páginas de front-end que somente admin possuem acesso
func AutenticaAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.ValidadeToken(c); err != nil {
			responses.ErrHTML(c, http.StatusBadRequest, "unauthorized.html", err)
			return
		}
		role, err := auth.PegarRoleUsuario(c)
		if err != nil {
			responses.ErrHTML(c, http.StatusBadRequest, "unauthorized.html", err)
			return
		}
		if role != "admin" {
			responses.ErrHTML(c, http.StatusForbidden, "unauthorized.html", ErrRoleInvalido)
			return
		}
		c.Next()
	}
}

func RespTypeJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
