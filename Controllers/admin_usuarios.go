package Controllers

import (
	"chat/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Home godoc
//
//	@Summary		Admin pegar usuario
//	@Description	é o admin para ser usado, apenas pelos admins, usado para ver usuarios.
//	@Tags			admin
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/admin/pegar_usuarios [get]
func Admin_usuarios(c *gin.Context) {
	rows ,err := Models.DB.Query(c.Request.Context(), "SELECT * FROM users")

	if err != nil{
		c.JSON(500, gin.H{
			"erro": "banco não encontrado",
		})
		return
	}

	var Mensagens []Models.Users

	for rows.Next(){
		var m Models.Users

		rows.Scan(&m.User_id,&m.Username,&m.Password_hash)
		Mensagens = append(Mensagens, m)
	}
	c.JSON(http.StatusOK, Mensagens)
}
