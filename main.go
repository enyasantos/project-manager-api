package main

import (
	"github.com/enyasantos/project-manager/config"
	"github.com/enyasantos/project-manager/server"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	server.Initialize()
}
