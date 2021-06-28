package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	ServerPort  int    `json:"server_port"`
	ServerMode  string `json:"server_mode"`
	DataRootDir string `json:"data_root_dir"`
	TempDir     string `json:"temp_dir"`
	DBHost      string `json:"db_host"`
	DBPort      uint   `json:"db_port"`
	DBUser      string `json:"db_user"`
	DBPwd       string `json:"db_pwd"`
	DBName      string `json:"db_name"`
	DBTimeZone  string `json:"db_timezone"`
}

var appConfig *AppConfig

func GetServerConfig() *AppConfig {
	if appConfig == nil {
		initServerConfig()
	}

	return appConfig
}

func initServerConfig() {
	jsonFile, err := os.Open("config/app.json")
	var config AppConfig

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[SUCCESS] open app config file")
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, &config)

	appConfig = &config
}

func (config *AppConfig) GetServerMode() string {
	switch mode := config.ServerMode; mode {
	case "debug":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
