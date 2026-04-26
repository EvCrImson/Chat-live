package Controllers

import (
	"chat/Models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Atualizar_mensagem(c *gin.Context) {
	tokenStr := Models.Pegar_Authorization(c)

	token, _ := Models.Validar_acess_token(tokenStr)

	dados := Models.Pegar_dados_de_acess_token(token)

	var mensagens Models.Mensagens_para_atualizar

	if erro := c.ShouldBindBodyWithJSON(&mensagens); erro != nil {
		c.JSON(400, gin.H{
			"erro": "Json invalido",
		})
		return
	}

	dados2, erro := strconv.Atoi(dados)

	if erro != nil {
		c.JSON(500, gin.H{
			"erro": "converção deu algum erro",
		})
		return
	}

	if mensagens.Mensagem_enviado_por != dados2 {
		c.JSON(400, gin.H{
			"erro": "usuario não pode atualizar essa mensagem",
		})
		return
	}

	_, err :=Models.DB.Exec(c.Request.Context(), "UPDATE mensagens SET mensagem = $1 WHERE id_mensagem = $2 AND mensagem = $3", mensagens.Mensagem_para_atulizar, mensagens.Id_mensagens, mensagens.Mensagem_antiga)

	if err != nil{
		c.JSON(400, gin.H{"error": "autuzalição da mensagem para o banco de dados, deu erro."})
		return
	}

	c.JSON(200, "mensagem atualizada")
}
