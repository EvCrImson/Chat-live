package Controllers

import (
	"chat/Models"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	if token == nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "token expirado",
			})
			return
		}
		c.AbortWithStatusJSON(401, gin.H{
			"error": "token invalido",
		})
		return
	}

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
