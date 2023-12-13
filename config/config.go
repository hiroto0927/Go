package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	UserName string
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
	dbConfig.UserName = os.Getenv("USER_NAME")
	dbConfig.Password = os.Getenv("PASSWORD")
	dbConfig.DBName = os.Getenv("DB")
	return dbConfig
}
