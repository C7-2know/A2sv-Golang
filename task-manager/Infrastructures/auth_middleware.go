package service

import (

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context){
		header,_:=c.Cookie("Authorization")
		err:=NewJwtService().ValidateToken(role,header)
		if err!=nil{
			c.JSON(401,gin.H{"error":err.Error()})
			c.Abort()
		}
		c.Next()
	}
}



