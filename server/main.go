package main

import (
	"fmt"
	"mini-oa-server/common/util/config"
	"mini-oa-server/service/users"

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

	v1 := r.Group("/v1")
	users_grp := v1.Group("/users")
	users.RegisterUserRouters(users_grp)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
