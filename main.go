package main

import (
	"os"

	"github/Babe-piya/book-collection/config"
	"github/Babe-piya/book-collection/server"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	cfg := config.LoadFileConfig(configPath)
	e, db := server.Start(cfg)

	server.Shutdown(e, db)
}
