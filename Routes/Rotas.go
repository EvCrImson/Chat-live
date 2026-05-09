package Routes

import (
	"chat/Controller"
	"chat/Middleware"
	"chat/Models"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

func Rotas() {
	r := gin.Default()
	r.GET(
		"/api/admin/pegar_usuarios",
		Middleware.AuthMiddlewareAdmin(), 
		Middleware.RateLimitByRoute(Models.Rdb, 1, time.Minute), 
		Controller.Admin_usuarios)
	r.GET(
		"/api/admin/pegar_conversas", 
		Middleware.AuthMiddlewareAdmin(), 
		Middleware.RateLimitByRoute(Models.Rdb, 1, time.Minute), 
		Controller.Admin_conversas)
	r.POST(
		"/api/criar_conversas", 
		Middleware.AuthMiddlewareNormal(), 
		Middleware.RateLimitByRoute(Models.Rdb, 15, time.Minute), 
		Controller.Criar_conversas)
	r.PUT(
		"/api/atualizar_mensagem", 
		Middleware.AuthMiddlewareNormal(), 
		Middleware.RateLimitByRoute(Models.Rdb, 5, time.Minute), 
		Controller.Atualizar_mensagem)
	r.DELETE(
		"/api/deletar_mensagem", 
		Middleware.AuthMiddlewareNormal(), 
		Middleware.RateLimitByRoute(Models.Rdb, 5, time.Minute),
		Controller.Deletar_mensagem)
	r.GET(
		"/api/pegar_conversas", 
		Middleware.AuthMiddlewareNormal(),
		Middleware.RateLimitByRoute(Models.Rdb, 10, time.Minute), 
		Controller.Ver_conversas)
	r.POST(
		"/api/cadrastar", 	
		Middleware.RateLimitByRoute(Models.Rdb, 5, time.Minute), 
		Controller.Criar_usuarios)
	r.DELETE(
		"/api/apagar_usuario", 
		Middleware.AuthMiddlewareNormal(), 
		Middleware.RateLimitByRoute(Models.Rdb, 1, time.Minute), 
		Controller.Apagar_usuario)
	r.GET(
		"/api/login", 
		Middleware.RateLimitByRoute(Models.Rdb, 4, time.Minute), 
		Controller.LoginUser)
	r.GET(
		"/api/criar_Acess_Token", 
		Middleware.RateLimitByRoute(Models.Rdb, 3, time.Minute*5), 
		Controller.Criar_Acess_token)
	r.Run(":" + os.Getenv("port"))
}
