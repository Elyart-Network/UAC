package main

import (
	"github.com/Elyart-Network/UAC/config"
	"github.com/Elyart-Network/UAC/data"
	"github.com/Elyart-Network/UAC/server"
)

func main() {
	config.EnvInit()
	data.Init()
	server.Start()
}
