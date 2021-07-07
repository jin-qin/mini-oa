package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRouters(r *gin.RouterGroup) {
	r.POST("/register", userRegisterRoute)
	r.POST("/login", userLoginRoute)
}

func userLoginRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func userRegisterRoute(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "register",
	})
}
