package Services

import (
	"chat/Models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Home godoc
//
//	@Summary		Logar
//	@Description	usuado para logar em usuarios, login.
//	@Tags			Usuario
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/login [post]
func LoginUser(c *gin.Context) {
	var usuario Models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": "json inválido"})
		return
	}

	var dadosVar Models.Dados
	err := Models.DB.QueryRow(c.Request.Context(), "SELECT password_hash,user_id FROM users WHERE username=$1", usuario.Name).Scan(&dadosVar.Password_hash, &dadosVar.User_id)

	if err != nil {
		c.JSON(401, gin.H{"error": "erro ao logar, teste trocar a senha ou o usuario"})
		return
	}

	erro := bcrypt.CompareHashAndPassword([]byte(dadosVar.Password_hash), []byte(usuario.Senha))

	if erro != nil {
		c.JSON(401, gin.H{"error": "erro ao logar, teste trocar a senha ou o usuario"})
		return
	}

	acess_token, _ := Models.Criar_acess_token(dadosVar.User_id)
	refresh_token, _ := Models.Criar_Refress_token(dadosVar.User_id)

	Models.Rdb.Set(c.Request.Context(), "refress:"+dadosVar.User_id, refresh_token, 7*24*time.Hour)

	c.JSON(201, gin.H{
		"refress_token": refresh_token,
		"acess_token":   acess_token,
	})
}
