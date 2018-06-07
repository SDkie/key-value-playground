package main

import (
	"fmt"
	"log"

	"github.com/SDkie/key-value-playground/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func DbInit() error {
	var err error
	cfg := config.GetConfig()
	dbCfg := &cfg.Database

	dbSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.DbName, dbCfg.Password)
	db, err = gorm.Open("postgres", dbSource)
	if err != nil {
		log.Println("Error connecting to db", err)
		return err
	}

	log.Println("Connected to db")
	return nil
}

func GetDb() *gorm.DB {
	return db
}
