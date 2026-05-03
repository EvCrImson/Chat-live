package Controller

import (
	"chat/Models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Home godoc
//
//	@Summary		Cadrastar
//	@Description	usuado para criar usuarios, cadrastar.
//	@Tags			Usuario
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/cadrastar [post]

func Criar_usuarios(c *gin.Context) {
	var usuario Models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao gerar hash"})
		return
	}

	_, err = Models.DB.Exec(c.Request.Context(),
		"INSERT INTO users (username, password_hash) VALUES ($1, $2)",
		usuario.Name,
		string(hash),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "Usuário criado com sucesso",
	})
}
