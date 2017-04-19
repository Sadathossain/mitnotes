package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/Sadathossain/mitnotes/notedb"
	"github.com/mcuadros/go-gin-prometheus"
)

var (
	appVersion  string
	showVersion bool
	database    notedb.NoteDB
)

func main() {
	configFile := flag.String("config-file", "./default.config", "Path to the configuration file")
	flag.BoolVar(&showVersion, "version", false, "Shows the version")
	flag.Parse()

	if showVersion {
		log.Printf("Version: %s\n", appVersion)
		return
	}

	config, err := readConfig(*configFile)
	if err != nil {
		log.Println(err)
	}

	gin.SetMode(config.ReleaseMode)
	if strings.ToLower(config.DBDriver) == "mysql" {
		database = notedb.NewMySQLDB(config.DBConfig, appVersion)
	} else if strings.ToLower(config.DBDriver) == "redis" {
		database = notedb.NewRedisDB(config.DBConfig, appVersion)
	}

	p := ginprometheus.NewPrometheus("gin")
	database.RegisterMetrics()

	// Iniitialize metrics
	quit := make(chan struct{})
	defer close(quit)
	if config.HealthCheckTime > 0 {
		healthCheckTimer := time.NewTicker(time.Duration(config.HealthCheckTime) * time.Second)
		go func() {
			for {
				select {
				case <-healthCheckTimer.C:
					log.Println("Called Health check")
					database.GetHealthStatus()
				case <-quit:
					healthCheckTimer.Stop()
					return
				}
			}
		}()
	}

	router := gin.Default()

	p.Use(router)
	router.GET("/read/note", readNoteHandler)
	router.GET("/insert/note/:value", insertNoteHandler)
	router.GET("/delete/note/:value", deleteNoteHandler)
	router.GET("/health", healthCheckHandler)
	router.GET("/whoami", whoAmIHandler)
	router.GET("/version", versionHandler)

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.Run(":3000")
}
