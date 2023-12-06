package app

import (
	"fmt"
	"mfahmii/golang-restful/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// func NewDB() *gorm.DB {
// 	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api")
// 	helper.PanicIfError(err)

// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	db.SetConnMaxIdleTime(10 * time.Minute)

// 	return db
// }

func OpenConnection(config *config.Config) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	dialect := mysql.Open(connectionString)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
