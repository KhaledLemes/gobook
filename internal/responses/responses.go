package responses

import (
	"encoding/json"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	if data != nil {
		c.Status(statusCode)
		if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Err(c *gin.Context, statusCode int, err error) {
	JSON(c, statusCode, struct {
		Error string `json:"error"`
	}{Error: primeiraToUpper(err.Error())})

}

func primeiraToUpper(s string) string {
	primeira, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(primeira)) + s[size:]
}
