package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbConfig := new(DBConfig)
	dbConfig.Host = os.Getenv("HOST")
	dbConfig.Port = os.Getenv("PORT")
	dbConfig.User = os.Getenv("USER")
	dbConfig.Password = os.Getenv("PASSWORD")
	dbConfig.DBName = os.Getenv("DB")
	return dbConfig
}
