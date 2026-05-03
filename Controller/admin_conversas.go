package Controller

import (
	"chat/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home godoc
//
//	@Summary		Admin pegar conversas
//	@Description	é o admin para ser usado, apenas pelos admins, usado para ver conversas.
//	@Tags			admin
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/admin/pegar_conversas [get]
func Admin_conversas(c *gin.Context) {
	rows, err := Models.DB.Query(c.Request.Context(), "SELECT * FROM mensagens")

	if err != nil {
		c.JSON(500, gin.H{
			"erro": "banco não encontrado",
		})
		return
	}

	var Mensagens []Models.Mensagens

	for rows.Next() {
		var m Models.Mensagens

		rows.Scan(&m.Id_mensagens, &m.Mensagem, &m.Mensagem_enviado_por, &m.Mensagem_recebida_por)
		Mensagens = append(Mensagens, m)
	}
	c.JSON(http.StatusOK, Mensagens)
}
