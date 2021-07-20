package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRouters(r *gin.RouterGroup) {
	r.GET("check_token", authCheckTokenRoute)
}

func authCheckTokenRoute(c *gin.Context) {
	c.Status(http.StatusOK)
}
