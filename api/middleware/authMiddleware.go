package middleware

import (
	"dependency_injection_tut/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


type AuthMiddleware struct{

}

func NewAuthMidleware() *AuthMiddleware{
	return &AuthMiddleware{}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func(am *AuthMiddleware) Authorized() gin.HandlerFunc{
return func(c *gin.Context){
	token:= ExtractToken(c)
	if (len([]rune(token))<1 ){
		c.JSON(http.StatusForbidden,gin.H{
			"error":"Unauthorized",
		})
		return
	}
	parsedToken,err:= utils.Validatetoken(token)
	if err!=nil{
	c.JSON(http.StatusForbidden,gin.H{
		"error":err.Error(),
	})
	}

	c.Set("user",gin.H{
		"parsedToken": &parsedToken.Claims,
	})
	c.Next()
}
}