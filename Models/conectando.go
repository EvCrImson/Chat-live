package Models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

var DB *pgxpool.Pool
var Rdb *redis.Client
var err error

func Connect_bancodedados(){
	conectbancdados := os.Getenv("dadosdebanco")

	DB, err = pgxpool.New(context.Background(), conectbancdados)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)	
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal("Banco não respondeu:", err)
	}

	fmt.Println("Conectado com sucesso")
}

func Connect_redis(){
	Rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("Addr"),
		Password: os.Getenv("Redis_senha"),
	})
}
