package server

import (
	"github.com/Elyart-Network/UAC/config"
	"github.com/Elyart-Network/UAC/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {

	// set log level
	logLevel, err := logrus.ParseLevel(config.Get("server.log_level").(string))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Start() {
	FileLogger := config.Get("server.file_logger").(bool)
	DebugMode := config.Get("server.debug_mode").(bool)
	ServerPort := config.Get("server.listen_port").(string)

	if FileLogger {
		gin.DisableConsoleColor()
		file, err := os.Create("app.log")
		if err != nil {
			log.Panicln(err)
		}
		gin.DefaultWriter = io.MultiWriter(file)
	}
	if DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LoggingMiddleware())
	Entry(engine)

	err := http.ListenAndServe(":"+ServerPort, engine)
	if err != nil {
		logrus.Fatalf("[Server] Failed to start server: %v", err)
	}
}
