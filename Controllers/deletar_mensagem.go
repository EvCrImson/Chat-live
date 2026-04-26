package Controllers

import (
	"chat/Models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Deletar_mensagem(c *gin.Context){
	tokenStr := Models.Pegar_Authorization(c)

	token, erro := Models.Validar_acess_token(tokenStr)

	if erro != nil{
		c.JSON(400, gin.H{
			"erro": "token invalido",
		})
		return
	}

	tokenStr = Models.Pegar_dados_de_acess_token(token)

	var mensagens Models.Mensagens
	if erro := c.ShouldBindBodyWithJSON(&mensagens); erro != nil{
		c.JSON(400, gin.H{
			"erro": "JSON invalido",
		})
		return
	}

	dados, _ := strconv.Atoi(tokenStr)

	if mensagens.Mensagem_enviado_por != dados{
		c.JSON(400, gin.H{
			"erro": "Mensagem não é sua",
		})
		return
	}

	_, err :=Models.DB.Exec(c.Request.Context(), "DELETE FROM mensagens WHERE id_mensagem = $1 AND mensagem = $2 AND mensagem_enviado_por = $3", mensagens.Id_mensagens, mensagens.Mensagem, mensagens.Mensagem_enviado_por)

	if err != nil{
		c.JSON(400, gin.H{"error": "autuzalição da mensagem para o banco de dados, deu erro."})
		return
	}

	c.JSON(200, "mensagem deletada")
}