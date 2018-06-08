package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SDkie/key-value-playground/db"
	"github.com/gorilla/websocket"
)

type Input struct {
	Key string `json:"key"`
}

type Output struct {
	Value string
}

type Keys struct {
	Id    int
	Key   string
	Value string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func KeyValue(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		input := Input{}
		err = json.Unmarshal(message, &input)
		if err != nil {
			log.Println("Error during Unmarshal", err)
			break
		}
		log.Println("Input:", input.Key)

		key := Keys{}
		err = db.GetDb().First(&key, "key = ?", input.Key).Error
		if err != nil {
			log.Println("Error during sql query", err)
			break
		}

		log.Println(key.Id, key.Key, key.Value)

		result := Output{}
		result.Value = key.Value

		outputMarshal, err := json.Marshal(result)
		if err != nil {
			log.Println("error marshal", err)
			break
		}

		err = c.WriteMessage(mt, outputMarshal)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}
