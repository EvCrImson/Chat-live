package Controllers

import (
	"chat/Models"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Home godoc
//
//	@Summary		Criar conversas
//	@Description	usuado para criar Conversas.
//	@Tags			Chat
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/criar_conversas [post]
func Criar_conversas(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if tokenStr == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "sem Authorization",
		})
		return
	}

	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "na Authorization sem o Bearer",
		})
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

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

	var mensagem_criar Models.Mensagem_para_criar

	if err := c.ShouldBindJSON(&mensagem_criar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dados Models.Users
	err = Models.DB.QueryRow(c.Request.Context(), "SELECT * FROM users WHERE user_id=$1", userid).Scan(&dados.User_id, &dados.Username, &dados.Password_hash)

	if err != nil || dados.Username == "" || dados.Password_hash == "" {
		c.JSON(400, gin.H{
			"erro": "usuario com esse user_id não achado",
		})
		return
	}

	_, err = Models.DB.Exec(c.Request.Context(), "INSERT INTO mensagens (mensagem, mensagem_enviado_por, mensagem_recebida_por) VALUES ($1, $2, $3)", mensagem_criar.Mensagem, userid, mensagem_criar.Mensagem_recebida_por)

	c.JSON(http.StatusOK, gin.H{
		"message": "hi",
	})
}
