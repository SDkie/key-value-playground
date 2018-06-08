package model

import (
	"log"

	"github.com/SDkie/key-value-playground/db"
)

type Keys struct {
	Id    int
	Key   string
	Value string
}

func GetValueFromKey(k string) (string, error) {
	key := Keys{}
	err := db.GetDb().First(&key, "key = ?", k).Error
	if err != nil {
		log.Println("Error during sql query", err)
		return "", err
	}
	return key.Value, nil
}
