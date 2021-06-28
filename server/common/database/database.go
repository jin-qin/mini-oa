package database

import (
	"fmt"
	"log"

	"mini-oa-server/common/util/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var dbInst *gorm.DB

// Using this function to get a connection.
func GetDB() *gorm.DB {
	if dbInst == nil {
		initDB()
	}
	return dbInst
}

// Opening a database and save the reference to `Database` struct.
func initDB() {
	conf := config.GetServerConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.DBHost,
		conf.DBUser,
		conf.DBPwd,
		conf.DBName,
		conf.DBPort,
		conf.DBTimeZone,
	)

	postgres_conf := postgres.Config{
		DSN: dsn,
	}

	db, err := gorm.Open(postgres.New(postgres_conf), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect database")
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	dbInst = db
}
