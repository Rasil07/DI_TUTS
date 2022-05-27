package utils

import (
	"dependency_injection_tut/model"
	"fmt"
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
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))

}

func Validatetoken(tokenString string) (*jwt.Token, error){
	token,err:= jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
		if _,ok:= token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,fmt.Errorf("unexpected sigining method: %v",token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")),nil

	})
	if err!=nil{
		return nil,err
	}
	
	return token,nil
}