package Models

import(
	"github.com/gin-gonic/gin"
	"strings"
)


func Pegar_Authorization(ctx *gin.Context) string{
	tokenStr := ctx.GetHeader("Authorization")

	if tokenStr == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "sem Authorization",
		})
		return "erro"
	}

	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "na Authorization sem o Bearer",
		})
		return "erro"
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	return tokenStr
}