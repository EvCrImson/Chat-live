package main

import (
	"chat/Middleware"
	"chat/Routes"
)

func main() {
	Middleware.Connecções()
	Routes.Rotas()
}
