package utils

import "os"

type DbConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

type ApiConfig struct {
	Host    string
	Port    string
	Address string
}

type Config struct {
	Db  *DbConfig
	Api *ApiConfig
}

func GetConfig() *Config {
	dbConfig := &DbConfig{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	apiConfig := &ApiConfig{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}

	return &Config{
		Db:  dbConfig,
		Api: apiConfig,
	}
}
