package controllers

import (
	"encoding/json"
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"gobook/internal/security"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CriaUsuario(c *gin.Context) {
	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(req, &usuario); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Prepara(true); err != nil {
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
	defer c.Request.Body.Close()

	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(req, &usuario); err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Prepara(false); err != nil {
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
	defer c.Request.Body.Close()

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
	ID := c.Request.URL.Query().Get("userID")

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
