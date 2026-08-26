package controllers

import (
	"errors"
	"fmt"
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func MostraTodasPropriedades(c *gin.Context) {
	var propriedades []models.Propriedade

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPropriedadesRepo(db)
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

	repo := repositories.NewPropriedadesRepo(db)

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

	repo := repositories.NewPropriedadesRepo(db)

	propriedade, err := repo.BuscaPropriedadePorID(ID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, propriedade)
}

func CriarPropriedade(c *gin.Context) {
	role, err := auth.ExtractUserRole(c)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	if role != "admin" && role != "owner" {
		responses.Err(c, http.StatusForbidden, errors.New("você não tem autorização para executar essa ação"))
	}

	var propriedade models.Propriedade
	if err := c.ShouldBindBodyWith(&propriedade, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userID, err := auth.ExtractUserID(c)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	fmt.Println(userID)

	repo := repositories.NewPropriedadesRepo(db)
	propriedadeNova, err := repo.CriaPropriedade(&propriedade, int(userID))
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "%s publicada!", propriedadeNova)
}

func EditarPropriedade(c *gin.Context) {

}

func DeletaPropriedadePorID(c *gin.Context) {

}
