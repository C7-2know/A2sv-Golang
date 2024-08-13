package service

import (

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context){
		header:=c.GetHeader("Authorization")
		err:=NewJwtService().ValidateToken(role,header)
		if err!=nil{
			c.Abort()
		}
		c.Next()
	}
}



