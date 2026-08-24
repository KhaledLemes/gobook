package controllers

import (
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MostraTodasPropriedas(c *gin.Context) {
	var propriedades []models.Propriedade

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsuariosRepo(db)
	propriedades, err = repo.BuscarTodasPropriedades()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, propriedades)
}

func BuscaPropriedadePorID(c *gin.Context) {

}

func CriarPropriedade(c *gin.Context) {

}

func EditarPropriedade(c *gin.Context) {

}

func DeletaPropriedadePorID(c *gin.Context) {

}
