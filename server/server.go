package server

import (
	"github.com/Elyart-Network/UAC/internal/actions"
	"github.com/Elyart-Network/UAC/internal/config"
	"github.com/Elyart-Network/UAC/internal/drivers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

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

	db := drivers.Database(config.Get("handlers.host").(string))
	actions.New(db)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	Entry(engine)

	err := http.ListenAndServe(":"+ServerPort, engine)
	if err != nil {
		log.Panicln(err)
	}
}