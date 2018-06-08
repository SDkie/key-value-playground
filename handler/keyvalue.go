package handler

import (
	"log"
	"net/http"

	"github.com/SDkie/key-value-playground/encoding"
	"github.com/SDkie/key-value-playground/model"
	"github.com/gorilla/websocket"
)

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

		input := encoding.Input{}
		err = input.ReadFromBytes(message)
		if err != nil {
			break
		}

		value, err := model.GetValueFromKey(input.Key)
		if err != nil {
			log.Println("Error getting value from Key", err)
			break
		}

		output := encoding.Output{
			Value: value,
		}

		outputMarshal, err := output.WriteToBytes()
		if err != nil {
			break
		}

		err = c.WriteMessage(mt, outputMarshal)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}
