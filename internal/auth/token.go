package auth

import (
	"errors"
	"fmt"
	"gobook/internal/config"
	"gobook/internal/models"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateToken criaa o token. Permissions é um objeto do tipo mapclaims, que tem o payload do JWT
// NewWithClaims efetivamente cria o token com a assinatura
func CreateToken(userID uint64, role models.Role) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Minute * 30).Unix()
	permissions["userID"] = userID
	permissions["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidadeToken(c *gin.Context) error {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func ExtractUserID(c *gin.Context) (int, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if permissions["userID"] == nil {
			return 0, errors.New(
				"userId está vazio",
			)
		}
		userID, err := strconv.Atoi(fmt.Sprintf("%.0f", permissions["userID"]))
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("token inválido")
}

func ExtractUserRole(c *gin.Context) (string, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if permissions["role"] == nil {
			return "", errors.New(
				"role está vazio",
			)
		}
		role := fmt.Sprintf("%s", permissions["role"])
		return role, nil
	}

	return "", errors.New("token inválido")
}

// extractToken extracts the token from header
func extractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnVerificationKey verifica o méthod de assinatura do token para ver se realmente é da família HMAC fazendo um Type Assertion
// Se conseguir fazer a conversão, significa que o method de assinatura está correto
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado. %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
