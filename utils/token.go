package utils

import (
	"dependency_injection_tut/model"
	"errors"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *model.User) (string, error) {

	token_lifespan,err := strconv.Atoi(os.Getenv("JWT_DURATION"))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func Validatetoken(tokenString string) (*jwt.Token, error){
	token,err:= jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")),nil
	})
	if err!=nil{
		return nil,err
	}
	if token.Valid{
		return token,nil
	}else{
		return nil,errors.New("invalid token")
	}
}