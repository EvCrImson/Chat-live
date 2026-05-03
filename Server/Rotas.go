package Server

import (
	"chat/Models"
	"chat/Services"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

func Rotas() {
	r := gin.Default()
	r.GET("/api/admin/pegar_usuarios", Models.AuthMiddlewareAdmin(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Services.Admin_usuarios)
	r.GET("/api/admin/pegar_conversas", Models.AuthMiddlewareAdmin(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Services.Admin_conversas)
	r.POST("/api/criar_conversas", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 15, time.Minute), Services.Criar_conversas)
	r.PUT("/api/atualizar_mensagem", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 5, time.Minute), Services.Atualizar_mensagem)
	r.DELETE("/api/deletar_mensagem", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 5, time.Minute), Services.Deletar_mensagem)
	r.GET("/api/pegar_conversas", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 10, time.Minute), Services.Ver_conversas)
	r.POST("/api/cadrastar", Models.RateLimitByRoute(Models.Rdb, 5, time.Minute), Services.Criar_usuarios)
	r.DELETE("/api/apagar_usuario", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Services.Apagar_usuario)
	r.GET("/api/login", Models.RateLimitByRoute(Models.Rdb, 4, time.Minute), Services.LoginUser)
	r.GET("/api/criar_Acess_Token", Models.RateLimitByRoute(Models.Rdb, 3, time.Minute*5), Services.Criar_Acess_token)
	r.Run(":" + os.Getenv("port"))
}
