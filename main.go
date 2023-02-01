package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"test-service-grpc/config"
	"test-service-grpc/server"
)

func main() {
	var cfg config.Config
	readConfig(&cfg)
	server.Start(&cfg)
}

func readConfig(cfg *config.Config) {
	configFileName := "config.json"
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}
