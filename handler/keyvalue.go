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

func (input *Input) ReadFromByte(msg []byte) error {
	err := json.Unmarshal(msg, input)
	if err != nil {
		log.Println("Error during Unmarshal", err)
		return err
	}
	return nil
}

type Output struct {
	Value string `json:"value"`
}

func (output *Output) WriteToByte() ([]byte, error) {
	outputMarshal, err := json.Marshal(*output)
	if err != nil {
		log.Println("Error during Marshal", err)
		return outputMarshal, err
	}

	return outputMarshal, nil
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
		err = input.ReadFromByte(message)
		if err != nil {
			break
		}

		value, err := model.GetValueFromKey(input.Key)
		if err != nil {
			log.Println("Error getting value from Key", err)
			break
		}

		output := Output{
			Value: value,
		}

		outputMarshal, err := output.WriteToByte()
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
