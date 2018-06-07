package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func DbInit() error {
	var err error

	db, err = gorm.Open("postgres", "host= port= user= dbname= password=")
	if err != nil {
		log.Println("Error connecting to db", err)
		return err
	}

	log.Println("Connected to db")
	return nil
}
