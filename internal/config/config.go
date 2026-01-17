package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	// struct tags
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:http_server`
}

func MainLoad() *Config {
	var configPath string
	// Config path by env
	configPath = os.Getenv("CONFIG_PATH")
	//Check config path
	if configPath == "" {
		// check in arguments
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path doesnot exist")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesnot exist: %s", err.Error())
	}
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read ConfigFile: %s", err.Error())
	}
	return &cfg
}
