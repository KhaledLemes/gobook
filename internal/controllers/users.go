package controllers

import (
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"gobook/internal/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CriaUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindWith(&usuario, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Prepara(true); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUsuariosRepo(db)

	usuario.ID, err = repo.CriarNovoUsuario(usuario)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindWith(&usuario, binding.JSON); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Prepara(false); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUsuariosRepo(db)
	userRetornado, err := repo.ProcurarPorEmail(usuario.Email)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	err = security.Compare(userRetornado.Senha, usuario.Senha)
	if err != nil {
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(uint64(userRetornado.ID), userRetornado.Role)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.Header("Authorization", token)
}

// BuscaPorID busca dados de um usuário pelo seu ID
// Usa o parâmetro userID na URL.
func BuscaPorID(c *gin.Context) {
	ID := c.Param("id")

	userID, err := strconv.Atoi(ID)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUsuariosRepo(db)

	usuario, err := repo.ProcuraPorID(userID)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, usuario)
}
