package handlers

import (
	"fmt"
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/handlers/erros"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"gobook/internal/security"
	"gobook/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CriaUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindWith(&usuario, binding.JSON); err != nil {
		if strings.Contains(err.Error(), "parsing time") {
			responses.Err(c, http.StatusBadRequest, erros.ErrDataNascInvalida)
			return
		}
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
		if utils.VerificaErro(err, "não existe no banco de dados") {
			responses.Err(c, http.StatusInternalServerError, erros.ErrCredenciaisInvalidas)
			return
		}
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	err = security.Compare(userRetornado.Senha, usuario.Senha)
	if err != nil {
		if utils.VerificaErro(err, "is not the hash of the given password") {
			responses.Err(c, http.StatusUnauthorized, erros.ErrCredenciaisInvalidas)
			return
		}
		responses.Err(c, http.StatusUnauthorized, err)
		return
	}
	usuario.Senha = ""

	fmt.Println(usuario.Nome)
	token, err := auth.CriarToken(uint64(userRetornado.ID), userRetornado.Role, userRetornado.Nome)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("auth", token, 0, "/", "", false, true)
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	c.SetCookie("auth", "out", -1, "/", "", false, true)
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

func Me(c *gin.Context) {
	nome, err := auth.PegarNomeUsuario(c)

	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	responses.JSON(c, http.StatusOK, struct {
		Nome string `json:"nome"`
	}{Nome: nome})
}
