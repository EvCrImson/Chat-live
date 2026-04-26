package Routes

import (
	"chat/Controllers"
	"chat/Models"
	_ "chat/docs"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/joho/godotenv"
)

func Rotas() {
	godotenv.Load() 
	Models.Connect_bancodedados()
	Models.Connect_redis()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/admin/pegar_usuarios", Models.AuthMiddlewareAdmin(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Controllers.Admin_usuarios)
	r.GET("/api/admin/pegar_conversas", Models.AuthMiddlewareAdmin(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Controllers.Admin_conversas)
	r.POST("/api/criar_conversas", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 15, time.Minute), Controllers.Criar_conversas)
	r.PUT("/api/atualizar_mensagem", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 5, time.Minute),  Controllers.Atualizar_mensagem)
	r.DELETE("/api/deletar_mensagem", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 5, time.Minute), Controllers.Deletar_mensagem)
	r.GET("/api/pegar_conversas", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 10, time.Minute), Controllers.Ver_conversas)
	r.POST("/api/cadrastar", Models.RateLimitByRoute(Models.Rdb, 5, time.Minute), Controllers.Criar_usuarios)
	r.DELETE("/api/apagar_usuario", Models.AuthMiddlewareNormal(), Models.RateLimitByRoute(Models.Rdb, 1, time.Minute), Controllers.Apagar_usuario)
	r.GET("/api/login", Models.RateLimitByRoute(Models.Rdb, 4, time.Minute), Controllers.Login)
	r.GET("/api/criar_Acess_Token", Models.RateLimitByRoute(Models.Rdb, 3, time.Minute*5), Controllers.Criar_Acess_token)
	r.Run(":" + os.Getenv("port"))
}
