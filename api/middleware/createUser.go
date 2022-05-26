package middleware

import (
	"dependency_injection_tut/api/controller"
	"dependency_injection_tut/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


type CreateUserPasswordMiddleware struct{} 



func NewCreateUserPasswordMiddleware() CreateUserPasswordMiddleware{
	return CreateUserPasswordMiddleware{}
}

func (cup CreateUserPasswordMiddleware) Handle() gin.HandlerFunc{
	return func(c *gin.Context){
		var input controller.CreateUserCredentials
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			}
			var passwordHash,err = utils.HashPassword(input.Password)
			if err!=nil{
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			}
			input.Password = passwordHash
			c.Set("payload",input)
			c.Next()
	}
}

