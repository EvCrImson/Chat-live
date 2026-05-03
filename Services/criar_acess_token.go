package Services

import (
	"chat/Models"

	"github.com/gin-gonic/gin"
)

// Home godoc
//
//	@Summary		criar acess token
//	@Description	é usado para criar acess token, usado em todos.
//	@Tags			Tokens
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/criar_Acess_Token [get]
func Criar_Acess_token(c *gin.Context) {

	var refresstoken Models.Refressrequest
	if err := c.ShouldBindJSON(&refresstoken); err != nil {
		c.JSON(400, gin.H{"error": "json inválido"})
		return
	}

	acessstoken, erro := Models.Validar_refress_token(refresstoken.Refresstoken)

	if erro != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "refress invalido",
		})
		return
	}

	user_id := Models.Pegar_dados_de_refresh_token(acessstoken)

	acess_token, erro := Models.Criar_acess_token(user_id)

	if erro != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "ao criar o acess token",
		})
		return
	}

	c.JSON(200, gin.H{
		"acess_token": acess_token,
	})
}
