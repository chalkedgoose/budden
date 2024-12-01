package config

import (
	"log"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port int `koanf:"port"`
	} `koanf:"server"`
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig() *Config {
	k := koanf.New(".")

	// Load JSON config file
	if err := k.Load(file.Provider("config.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load environment variables
	if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
		log.Fatalf("error loading env vars: %v", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	return &cfg
}
