package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config model
type Config struct {
	Live    bool   `json:"live"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	Rethink struct {
		Host    string `json:"host"`
		Db      string `json:"db"`
		MaxOpen int    `json:"max_open"`
	}
	UploadPath string `json:"upload_path"`
	Mail       struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}
}

const configPath = "config.json"

var c Config

func init() {
	log.Println("[config]", "loading...")

	load()

	log.Println("[config]", "loaded!")
}

// Get configuration
func Get() Config {
	return c
}

func load() {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("[config]", err)
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&c); err != nil {
		log.Fatalln("[config]", err)
	}
}
