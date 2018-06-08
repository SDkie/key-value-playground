package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SDkie/key-value-playground/model"
	"github.com/gorilla/websocket"
)

type Input struct {
	Key string `json:"key"`
}

type Output struct {
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

		value, err := model.GetValueFromKey(input.Key)
		if err != nil {
			log.Println("Error getting value from Key", err)
			break
		}

		result := Output{
			Value: value,
		}

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
