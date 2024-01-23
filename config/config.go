package config

import (
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host     string `default:"localhost"`
	Port     string
	Dbname   string
	Username string
	Password string
}

type Config struct {
	Database DBConfig
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	AppConfig = &Config{}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	AppConfig.Database = dbConfig
}
