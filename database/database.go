package database

import (
	"myapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	DB := config.NewDBConfig()

	dsn := "host=" + DB.Host +
		" user=" + DB.UserName +
		" password=" + DB.Password +
		" dbname=" + DB.DBName +
		" port=" + DB.Port +
		" sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
