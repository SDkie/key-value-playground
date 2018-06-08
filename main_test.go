package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/SDkie/key-value-playground/config"
	"github.com/SDkie/key-value-playground/db"
	"github.com/SDkie/key-value-playground/encoding"
	"github.com/SDkie/key-value-playground/handler"
	"github.com/gorilla/websocket"
)

func TestMain(m *testing.M) {
	config.Init()

	err := db.Init()
	if err != nil {
		log.Panic(err)
	}

	testResult := m.Run()

	db.Close()
	os.Exit(testResult)
}

func TestKeyValueHandler(t *testing.T) {
	cases := []struct {
		Key, Value string
	}{
		{"0", "零"},
		{"1", "Uno"},
		{"2", "Two"},
		{"3", "तीन"},
		{"4", "أربعة"},
		{"5", "Cinco"},
		{"6", "ছয়"},
		{"7", "семь"},
		{"8", "八"},
		//		{"9", "ਨੌਂ"},
	}

	s := httptest.NewServer(http.HandlerFunc(handler.KeyValue))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	for _, c := range cases {
		var output encoding.Output
		input := encoding.Input{Key: c.Key}
		inputMarshal, err := input.WriteToByte()
		if err != nil {
			t.Fatal("Error during Marshal:", err)
		}

		if err := ws.WriteMessage(websocket.TextMessage, inputMarshal); err != nil {
			t.Fatal("Error writing on WebSocket:", err)
		}

		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatal("Error Reading Message:", err)
		}

		err = output.ReadFromByte(p)
		if err != nil {
			t.Fatalf("Error During UnMarshal:", err)
		}

		if c.Value != output.Value {
			t.Fatalf("Not Match for Key:%s. Expected:%s Value returned:%s", c.Key, c.Value, output.Value)
		}
	}

	// Close WebSocket
	err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		t.Fatal("Error while closing WebSocket:", err)
		return
	}

}
