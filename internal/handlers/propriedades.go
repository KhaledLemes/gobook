package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gobook/internal/auth"
	"gobook/internal/database"
	"gobook/internal/models"
	"gobook/internal/repositories"
	"gobook/internal/responses"
	"gobook/utils"
	"net/http"
	"path"
	"strconv"
	"strings"

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

func BuscarOitoPrimeirasPropriedadesAleatorio(c *gin.Context) {
	var propriedades []models.Propriedade

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPropriedadesRepo(db)
	propriedades, err = repo.BuscaroITOPrimeirasPropriedadesAleatorio()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, propriedades)
}

func BuscarTodasPorDono(c *gin.Context) {
	var propriedades []models.Propriedade

	db, err := database.Connect()
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusForbidden, err)
		return
	}

	repo := repositories.NewPropriedadesRepo(db)
	propriedades, err = repo.BuscarTodasPorDono(userID)
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
	role, err := auth.PegarRoleUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	if role != "admin" && role != "owner" {
		responses.Err(c, http.StatusForbidden, errors.New("você não tem autorização para executar essa ação"))
	}

	var propriedade models.Propriedade

	arquivo, err := c.FormFile("img")
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}
	data := c.PostForm("data")

	if err := json.Unmarshal([]byte(data), &propriedade); err != nil {
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

	userID, err := auth.PegarIDUsuario(c)
	if err != nil {
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	bytes := make([]byte, 16)
	_, err = rand.Read(bytes)
	if err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}
	randstr := hex.EncodeToString(bytes)
	extensao := strings.Split(arquivo.Filename, ".")[1]
	nomeDoArquivo := fmt.Sprintf("%s.%s", randstr, extensao)

	repo := repositories.NewPropriedadesRepo(db)
	_, err = repo.CriaPropriedade(&propriedade, int(userID), nomeDoArquivo)
	if err != nil {
		if utils.VerificaErro(err, "Duplicate entry") {
			msg := fmt.Sprintf("propriedade chamada %s já existe", propriedade.Nome)
			responses.Err(c, http.StatusBadRequest, errors.New(msg))
			return
		}
		responses.Err(c, http.StatusBadRequest, err)
		return
	}

	caminho := path.Join("./web/static/img/propriedades", nomeDoArquivo)
	fmt.Println(caminho)
	if err = c.SaveUploadedFile(arquivo, caminho); err != nil {
		responses.Err(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
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
