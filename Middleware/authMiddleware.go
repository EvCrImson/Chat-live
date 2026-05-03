package Middleware

import (
	"errors"
	"strings"
	"chat/Models"
	"chat/Services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddlewareNormal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := Models.Pegar_Authorization(ctx)

		_, err := Services.Validar_acess_token(tokenStr)

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.AbortWithStatusJSON(401, gin.H{
					"error": "token expirado",
				})
				return
			}
			ctx.AbortWithStatusJSON(401, gin.H{ 	
				"error": "token invalido",
			})
			return
		}

		ctx.Next()
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "sem Authorization",
			})
			return
		}

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "na Authorization sem o Bearer",
			})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		token, err := Services.Validar_acess_token(tokenStr)

		if token == nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.AbortWithStatusJSON(401, gin.H{
					"error": "token expirado",
				})
				return
			}
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		userid := Services.Pegar_dados_de_acess_token(token)

		if userid != "1" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "user_id errado ou o role",
			})
			return
		}

		ctx.Next()
	}
}
