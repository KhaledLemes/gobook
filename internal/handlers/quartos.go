package handlers

import (
	"database/sql"
	"errors"
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

func BuscarTodosQuartosDoDono(c *gin.Context) {
	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	quartos, err := repositories.NewQuartosRepo(db).BuscarTodosQuartosDeUsuarioPorID(userID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, quartos)
}

func BuscarTodosQuartosDaPropriedade(c *gin.Context) {
	propriedadeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.Err(c, http.StatusBadRequest, errors.New("id da propriedade não informado"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	quartos, err := repositories.NewQuartosRepo(db).BuscarTodosQuartosDePropriedadePorID(propriedadeID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, quartos)
}

func BuscarQuartoPorNome(c *gin.Context) {
	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	quarto, err := repositories.NewQuartosRepo(db).BuscaQuartoPorNome(c.Param("nome"), userID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}
		responses.Err(c, status, err)
		return
	}

	c.IndentedJSON(http.StatusOK, quarto)
}

func CriarQuarto(c *gin.Context) {
	role, err := auth.PegarRoleUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}
	if role != "owner" && role != "admin" {
		responses.Err(c, http.StatusForbidden, errors.New("você não tem autorização para executar essa ação"))
		return
	}

	var quarto models.Quarto
	if err := c.ShouldBindWith(&quarto, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	propID, err := strconv.Atoi(c.Query("propID"))
	if err != nil {
		responses.Err(c, http.StatusBadRequest, errors.New("id da propriedade inválido"))
		return
	}
	quarto.Propriedade = propID
	if err := quarto.Validar(); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Arrumar depois e arrumar um modo melhor de fazer essa verificação. Isso pode quebrar em ambientes concorrentes.
	if _, err := repositories.NewPropriedadesRepo(db).VerificaDono(quarto.Propriedade, userID); err != nil {
		responses.Err(c, http.StatusForbidden, err)
		return
	}

	quarto.Dono = userID
	if err := repositories.NewQuartosRepo(db).CriaQuarto(quarto, userID); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}

func AtualizarQuarto(c *gin.Context) {
	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	quartoID, err := strconv.Atoi(c.Param("id"))
	if err != nil || quartoID <= 0 {
		responses.Err(c, http.StatusBadRequest, errors.New("id do quarto inválido"))
		return
	}

	var quarto models.Quarto
	if err := c.ShouldBindWith(&quarto, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	if err := quarto.Validar(); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	quarto.ID = quartoID
	quarto.Dono = userID

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if err := repositories.NewQuartosRepo(db).AtualizaQuarto(quarto); err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func DeletarQuartoPorNome(c *gin.Context) {
	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if err := repositories.NewQuartosRepo(db).DeletaQuartoPorNome(c.Param("nome"), userID); err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
