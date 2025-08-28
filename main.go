package main

import (
	"log"

	"gifma-backend/config"
	"gifma-backend/internal/api"
)

func main() {
	cfg, err := config.SetupEnv(".")

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
