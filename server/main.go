package main

import (
	"fmt"
	"mini-oa-server/common/util/config"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := config.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
