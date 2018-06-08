package model

import (
	"errors"
	"log"

	"github.com/SDkie/key-value-playground/db"
	"github.com/jinzhu/gorm"
)

type Keys struct {
	Id    int
	Key   string
	Value string
}

func GetValueFromKey(k string) (string, error) {
	key := Keys{}
	err := db.GetDb().First(&key, "key = ?", k).Error
	if err == gorm.ErrRecordNotFound {
		return "", errors.New("Invalid Key")
	} else if err != nil {
		log.Println("Error during sql query", err)
		return "", err
	}
	return key.Value, nil
}
