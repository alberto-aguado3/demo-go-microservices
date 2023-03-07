package configuration

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Sql SqlConfig `envPrefix:"DB"`

	GinMode string `env:"MODE" default:"debug"`
}

type SqlConfig struct {
	User     string        `env:"USER"` //struct field tag
	Password string        `env:"PASSWORD"`
	Host     string        `env:"HOST"`
	Port     string        `env:"PORT"`
	DbName   string        `env:"NAME"`
	Timeout  time.Duration `env:"TIMEOUT"`
}

// ReadEnv - Load env variables from ".env" file, and return a Configuration with the env vars assigned to it.
func ReadEnv() (*Configuration, error) {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error reading ../.end")
		if err := godotenv.Load("./.env"); err != nil {
			fmt.Println("Error reading ./.end")
			return nil, err
		}
	}

	options := env.Options{
		Prefix: "API_",
	}

	var cfg Configuration
	err := env.Parse(&cfg, options)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Configuration:\n%+v\n", cfg)
	return &cfg, nil
}
