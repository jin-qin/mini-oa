package main

import (
	"fmt"
	"mini-oa-server/common/util/config"
	"mini-oa-server/service/auth"
	"mini-oa-server/service/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AutoMigrate() {
	users.AutoMigrate()
}

func main() {
	AutoMigrate()

	appConfig := config.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Enable CORS for all origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.Use(static.Serve("/app", static.LocalFile("./client", false)))

	v1 := r.Group("/v1")
	users_grp := v1.Group("/users")
	users.RegisterUserRouters(users_grp)

	v1.Use(users.AuthMiddleware(true))

	auth_grp := v1.Group("auth")
	auth.RegisterAuthRouters(auth_grp)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
