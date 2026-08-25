package controllers

import (
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MostraTodasPropriedades(c *gin.Context) {
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

func BuscaPropriedadePorNome(c *gin.Context) {
	nome := c.Param("nome")

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsuariosRepo(db)

	propriedade, err := repo.BuscaPropriedadePorNome(nome)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, propriedade)
}

func BuscaPropriedadePorID(c *gin.Context) {
	ID := c.Param("id")

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsuariosRepo(db)

	propriedade, err := repo.BuscaPropriedadePorID(ID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, propriedade)
}

func CriarPropriedade(c *gin.Context) {

}

func EditarPropriedade(c *gin.Context) {

}

func DeletaPropriedadePorID(c *gin.Context) {

}
