package main

import (
	"encoding/json"
	"fmt"
	"time"

	"./src"

	"github.com/gorilla/websocket"
)

type generic struct {
	Type        string `json:"type"`
	JSONString  string `json:"jsonstring"`
	JSONString2 string `json:"jsonstring2"`
}

var host = "localhost"
var port = "8080"
var isMining = false
var genericObject generic

var currentBlock structs.Block
var currentBuffer structs.Buffer

func connectToServer(conn *websocket.Conn) {
	var result generic
	result.Type = "connecting"
	fmt.Println("Connecting to: " + host + ":" + port + ".")
	conn.WriteMessage(websocket.TextMessage, genericMarshall(result))
}

func main() {
	URL := "ws://" + host + ":" + port + ""

	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	connectToServer(conn)
	go iter(conn)
	for {
		time.Sleep(time.Millisecond * 100)
	}
}

func iter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Millisecond * 400)
		var incomming generic
		err := conn.ReadJSON(&incomming)
		if err != nil {
			fmt.Println(err)
			return
		}
		decode(incomming, conn)
	}
}

func decode(incommingMessage generic, conn *websocket.Conn) {

	genericObject.Type = ""
	genericObject.JSONString = ""

	if incommingMessage.Type == "connected" {
		genericObject.Type = "connected"
		fmt.Println("Connected to server.")
		conn.WriteMessage(websocket.TextMessage, genericMarshall(genericObject))
	} else if incommingMessage.Type == "startWorker" {
		err := json.Unmarshal([]byte(incommingMessage.JSONString), &currentBlock)
		if err != nil {
			println(err)
		}

		err2 := json.Unmarshal([]byte(incommingMessage.JSONString2), &currentBuffer)
		if err2 != nil {
			println(err2)
		}
		if isMining {
			println("Head changed, recieving new block!")
		} else {
			isMining = true
			println("Start mining!")
		}
	} else if incommingMessage.Type == "stopWorker" {
		isMining = false
		println("Stop mining!")
	} else {
		genericObject.Type = "ping"
		genericObject.JSONString = ""
		conn.WriteMessage(websocket.PingMessage, genericMarshall(genericObject))
	}
}

func genericMarshall(generic generic) []byte {
	result, _ := json.Marshal(generic)
	return result
}
