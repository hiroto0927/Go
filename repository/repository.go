package repository

import (
	"myapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := "host=" + config.GetDBConfig().Host +
		" user=" + config.GetDBConfig().User +
		" password=" + config.GetDBConfig().Password +
		" dbname=" + config.GetDBConfig().DBName +
		" port=" + config.GetDBConfig().Port +
		" sslmode=" + config.GetDBConfig().SSLMode +
		" TimeZone=" + config.GetDBConfig().TimeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
