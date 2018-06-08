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

	var message, outputBytes []byte
	var value string

	for {
		if err != nil {
			err = c.WriteMessage(websocket.TextMessage, []byte("ERROR:"+err.Error()))
			if err != nil {
				log.Println("Error while sending msg,", err)
				continue
			}
		}

		_, message, err = c.ReadMessage()
		if err != nil {
			log.Println("Error while reading msg,", err)
			continue
		}

		input := encoding.Input{}
		err = input.ReadFromBytes(message)
		if err != nil {
			continue
		}

		value, err = model.GetValueFromKey(input.Key)
		if err != nil {
			log.Println("Error getting value from Key", err)
			continue
		}

		output := encoding.Output{
			Value: value,
		}

		outputBytes, err = output.WriteToBytes()
		if err != nil {
			continue
		}

		err = c.WriteMessage(websocket.TextMessage, outputBytes)
		if err != nil {
			log.Println("Error while Writing msg:", err)
			continue
		}
	}
}
