package config

import (
	"os"
)

// DB接続情報
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
	TimeZone string
}

// DB接続情報を取得
func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		DBName:   os.Getenv("DB"),
		Password: os.Getenv("PASSWORD"),
		SSLMode:  os.Getenv("SSL_MODE"),
		TimeZone: os.Getenv("TIME_ZONE"),
	}
}
