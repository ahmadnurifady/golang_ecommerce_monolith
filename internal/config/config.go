package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
	Timezone string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}

	if c.ApiPort == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" ||
		c.DbConfig.Password == "" || c.Driver == "" {
		return errors.New("missing required environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
