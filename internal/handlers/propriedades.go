package handlers

import (
	"errors"
	"fmt"
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"net/http"
	"strconv"

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

	if err = propriedade.Prepara(); err != nil {
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

	repo := repositories.NewPropriedadesRepo(db)
	propriedadeNova, err := repo.CriaPropriedade(&propriedade, int(userID))
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "A propriedade %s foi publicada!", propriedadeNova)
}

func EditarPropriedade(c *gin.Context) {
	var propriedade models.Propriedade
	userID, admin, err := auth.VerificaOwnerEAdmin(c)
	if err != nil {
		responses.Err(c, http.StatusForbidden, err)
		return
	}

	strPropID := c.Param("id")
	propID, err := strconv.Atoi(strPropID)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindWith(&propriedade, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	if err := propriedade.Prepara(); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	fmt.Println(userID)
	repo := repositories.NewPropriedadesRepo(db)
	if !admin {
		_, err = repo.VerificaDono(propID, userID)
		if err != nil {
			responses.Err(c, http.StatusForbidden, err)
			return
		}
	}
	if err = repo.EditaPropriedade(propriedade, propID); err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, "Sucesso na edição!")
}

func DeletaPropriedadePorID(c *gin.Context) {
	userID, admin, err := auth.VerificaOwnerEAdmin(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	strPropID := c.Param("id")
	propID, err := strconv.Atoi(strPropID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPropriedadesRepo(db)

	// Verifica se é realmente o dono tentando apagar se não for um admin
	var propriedadeNome string
	if !admin {
		propriedadeNome, err = repo.VerificaDono(propID, userID)
		if err != nil {
			responses.Err(c, http.StatusForbidden, err)
			return
		}
	} else {
		propriedadeNome, err = repo.RetornaNome(propID)
		if err != nil {
			responses.Err(c, http.StatusForbidden, err)
			return
		}
	}

	if err = repo.DeletaPropriedade(propID); err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "A propriedade %s foi deletada com sucesso!", propriedadeNome)
}
