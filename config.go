package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type NoteAppConfig struct {
	HealthCheckTime int
	DBDriver        string
	DBConfig        map[string]string
	ReleaseMode     string
}

func readConfig(configFile string) (*NoteAppConfig, error) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return &NoteAppConfig{
			DBDriver:    "redis",
			DBConfig:    map[string]string{},
			ReleaseMode: gin.DebugMode,
		}, err
	}
	config := &NoteAppConfig{}
	json.Unmarshal(file, config)

	if config.DBDriver == "" {
		log.Println("Use redis as default")
		config.DBDriver = "redis"
	}

	if config.DBConfig == nil {
		config.DBConfig = map[string]string{}
	}

	if config.ReleaseMode == "" {
		config.ReleaseMode = gin.DebugMode
	}

	return config, err
}
