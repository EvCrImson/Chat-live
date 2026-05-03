package Services

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var refreshSecret = []byte(os.Getenv("Refress_token_chave"))

func Criar_Refress_token(user_id string) (string, error){	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"User_id": user_id, "exp": time.Now().Add(7* 24 * time.Hour).Unix()})

	return token.SignedString(refreshSecret)
}

func Validar_refress_token(refress_token string) (*jwt.Token, error){
	token, err := jwt.Parse(refress_token, func(t *jwt.Token) (interface{}, error) {return refreshSecret, nil})
	
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return token, nil
}

func Pegar_dados_de_refresh_token(refreshtoken *jwt.Token) (string){
	claims := refreshtoken.Claims.(jwt.MapClaims)
	userID := claims["User_id"].(string)

	return userID
}