package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

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

func KeyValue(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

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
		err = db.First(&key, "key = ?", input.Key).Error
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

func main() {
	DbInit()
	http.HandleFunc("/keyvalue", KeyValue)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
