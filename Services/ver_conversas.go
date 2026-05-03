package Services

import (
	"chat/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Home godoc
//
//	@Summary		ver conversas
//	@Description	usuado para ver Conversas.
//	@Tags			Chat
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/pegar_conversas [get]
func Ver_conversas(c *gin.Context) {
	tokenStr := Models.Pegar_Authorization(c)

	token, err := Models.Validar_acess_token(tokenStr)

	userid := Models.Pegar_dados_de_acess_token(token)

	rows, err := Models.DB.Query(c.Request.Context(), "SELECT mensagem,mensagem_recebida_por FROM mensagens WHERE mensagem_enviado_por=$1", userid)

	var mensagens []Models.Mensagens

	if err != nil {
		c.JSON(400, gin.H{
			"error": "não achado no banco de dados, nenhuma mensagem",
		})
		return
	}

	for rows.Next() {
		var m Models.Mensagens

		rows.Scan(&m.Mensagem, &m.Mensagem_recebida_por)
		mensagens = append(mensagens, m)
	}

	c.JSON(http.StatusOK, mensagens)
}
