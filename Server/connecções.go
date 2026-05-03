package Server

import (
	"chat/Models"
	"github.com/joho/godotenv"
)

func Connecções(){
	godotenv.Load()
	Models.Connect_bancodedados()
	Models.Connect_redis()
}