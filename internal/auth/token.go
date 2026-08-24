package auth

import (
	"errors"
	"fmt"
	"gobook/internal/config"
	"gobook/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

func ValidadeToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("PARA DEBUG, DENTRO DE ExtractUserID:", permissions["userID"])
		if permissions["userID"] == nil {
			return 0, errors.New(
				"UserID is a <nil> value. Please, revaluate your trash code or choose another career that does not demand writing code",
			)
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid token")
}

// extractToken extracts the token from header
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnVerificationKey verifica o méthod de assinatura do token para ver se realmente é da família HMAC fazendo um Type Assertion
// Se conseguir fazer a conversão, significa que o method de assinatura está correto
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method. %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
