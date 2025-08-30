package config

import (
	// "fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type APP struct {
	Env string `yaml:"env" env:"APP_ENV" env-default:"dev"`
}
type HTTP struct {
	Address           string        `yaml:"address" env:"HTTP_ADDRESS" env-default:":8080"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"HTTP_READ_HEADER_TIMEOUT" env-default:"5s"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT" env-default:"10s"`
}
type CORS struct {
	AllowedOrigins []string `yaml:"allowed_origins" env:"CORS_ALLOWED_ORIGINS"  env-separator:"," env-default:"*"`
}

type Config struct {
	APP  APP  `yaml:"app"`
	HTTP HTTP `yaml:"http"`
	CORS CORS `yaml:"cors"`
}

func MustLoad() *Config {

	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "local.yaml"
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to load config: %s: %v", configPath, err)
	}

	//checking all the  critical fields
	//address

	if _, err := net.ResolveTCPAddr("tcp", cfg.HTTP.Address); err != nil {
		log.Fatalf("Failed to resolve HTTP address %q: %v", cfg.HTTP.Address, err)
	}

	if cfg.HTTP.ReadHeaderTimeout <= 0 {
		log.Fatalf("Time out cant be zero, but got this %v", cfg.HTTP.ReadHeaderTimeout)
	}

	if cfg.HTTP.ShutdownTimeout <= 0 {
		log.Fatalf("Time out cant be zero, but got this %v", cfg.HTTP.ShutdownTimeout)
	}

	if len(cfg.CORS.AllowedOrigins) == 0 {
		log.Fatalf("CORS.AllowedOrigins cannot be empty")
	}

	// fmt.Printf("config %+v", cfg)

	return &cfg
}
