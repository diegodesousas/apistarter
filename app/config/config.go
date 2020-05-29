package config

import (
	"log"

	"github.com/spf13/viper"
)

type Database struct {
	Driver string `mapstructure:"database_driver"`
	DSN    string `mapstructure:"database_dsn"`
}

type Config struct {
	Database Database `mapstructure:",squash"`
}

func New() (*Config, error) {
	viper.SetDefault("DATABASE_DRIVER", "postgres")
	viper.SetDefault("DATABASE_DSN", "postgres://postgres:root@postgres11.hud:5432/apistarter?sslmode=disable")

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config: .env file not found")
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
