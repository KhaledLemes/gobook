package responses

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	if data != nil {
		if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Err(c *gin.Context, statusCode int, err error) {
	c.Status(statusCode)
	JSON(c, statusCode, struct {
		Error string `json:"error"`
	}{Error: err.Error()})

}
