package Models

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var accessSecret = []byte(os.Getenv("Acess_token_chave"))
var refreshSecret = []byte(os.Getenv("Refress_token_chave"))

func Criar_acess_token(user_id string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"User_id": user_id, "exp": time.Now().Add(15 * time.Minute).Unix()})

	return token.SignedString(accessSecret)
}


func Criar_Refress_token(user_id string) (string, error){	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"User_id": user_id, "exp": time.Now().Add(7* 24 * time.Hour).Unix()})

	return token.SignedString(refreshSecret)
}

func Atualizar_acess_token(refreshtoken *jwt.Token) (string, error){
	claims := refreshtoken.Claims.(jwt.MapClaims)
	userID := claims["User_id"].(string)

	return Criar_acess_token(userID)
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

func Validar_acess_token(acesstoken string) (*jwt.Token, error){
	token, err := jwt.Parse(acesstoken, func(t *jwt.Token) (interface{}, error) {return accessSecret, nil})
	
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}
	return token, nil
}

func Pegar_dados_de_acess_token(acesstoken *jwt.Token) (string){
	claims := acesstoken.Claims.(jwt.MapClaims)
	userID := claims["User_id"].(string)

	return userID
}

func Pegar_dados_de_refresh_token(refreshtoken *jwt.Token) (string){
	claims := refreshtoken.Claims.(jwt.MapClaims)
	userID := claims["User_id"].(string)

	return userID
}