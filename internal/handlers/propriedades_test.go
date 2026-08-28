package handlers

import (
	"encoding/json"
	"fmt"
	"gobook/internal/config"
	"gobook/internal/models"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMostraTodasPropriedades(t *testing.T) {
	// Recorder é basicamente quem vai receber as respostas
	recorder := httptest.NewRecorder()
	// Nosso contexto forjado do zero
	c := GetTestGinContext(recorder)

	// Têm que existir
	params := []gin.Param{
		{
			Key:   "d",
			Value: "a",
		},
	}
	u := url.Values{}

	Make(c, http.MethodGet, params, u)

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config.Load()
	MostraTodasPropriedades(c)

	var propriedades []models.Propriedade
	if ok := assert.EqualValues(t, http.StatusOK, recorder.Code); !ok {
		t.Errorf("Queria %d mas recebi %d\n", http.StatusForbidden, recorder.Code)
	}
	propriedadesStr := recorder.Body.Bytes()
	json.Unmarshal(propriedadesStr, &propriedades)
	for i, p := range propriedades {
		fmt.Println("Propriedade ", i+1, ": \n", p)
	}
}

func TestBuscaPropriedadePorNome(t *testing.T) {
	recorder := httptest.NewRecorder()
	c := GetTestGinContext(recorder)

	params := []gin.Param{
		{
			Key:   "nome",
			Value: "naoexiste",
		},
		{
			Key:   "nome",
			Value: "shopping_penha",
		},
	}
	u := url.Values{}

	for i, _ := range params {
		Make(c, http.MethodGet, params, u)
		if i == 0 {
			BuscaPropriedadePorID(c)
			if !assert.EqualValues(t, http.StatusOK, recorder.Code) {
				t.Errorf("esperado:")
			}
		}

	}
}

// Primeiro criamos o nosso Contexto para usar no teste
func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(recorder)

	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return c
}

func Make(c *gin.Context, method string, params gin.Params, u url.Values) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "applications/json")

	c.Params = params
	c.Request.URL.RawQuery = u.Encode()
}
