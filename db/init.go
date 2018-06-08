package db

import (
	"fmt"
	"log"

	"github.com/SDkie/key-value-playground/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() error {
	var err error
	cfg := config.GetConfig()
	dbCfg := &cfg.Database

	dbSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbCfg.DB_HOST, dbCfg.DB_PORT, dbCfg.DB_USER, dbCfg.DB_NAME, dbCfg.DB_PASSWORD)
	db, err = gorm.Open("postgres", dbSource)
	if err != nil {
		log.Println("Error connecting to db", err)
		return err
	}

	log.Println("Database Initialized")
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func Close() error {
	err := db.Close()
	if err != nil {
		log.Println("Error while closing Database", err)
	}

	return err
}
