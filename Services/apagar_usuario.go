package Services

import (
	"chat/Models"

	"github.com/gin-gonic/gin"
)

func Apagar_usuario(c *gin.Context) {
	tokenStr := Models.Pegar_Authorization(c)
	token, _ := Models.Validar_acess_token(tokenStr)
	tokenStr = Models.Pegar_dados_de_acess_token(token)

	var user Models.Deletar_usuario

	if erro := c.ShouldBindBodyWithJSON(&user); erro != nil {
		c.JSON(400, gin.H{
			"erro": "JSON invalido",
		})
		return
	}

	if user.User_id != tokenStr {
		c.JSON(400, gin.H{
			"erro": "você não pode deletar um usuario que não é o seu",
		})
		return
	}
	_, err := Models.DB.Exec(c.Request.Context(), "DELETE FROM users WHERE user_id = $1 AND username = $2", user.User_id, user.Usuario)

	if err != nil {
		c.JSON(500, gin.H{
			"erro": "aconteceu algum erro, desculpe",
		})
		return
	}

	c.JSON(200, "usuario deletado com sucesso")
}
