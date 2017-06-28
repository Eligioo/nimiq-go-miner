package main

import (
	"encoding/json"
	"fmt"
	"time"

	"./src/structs"

	"strconv"

	"github.com/gorilla/websocket"
)

var host = "localhost"
var port = "8080"
var isMining = false
var isConnected = false
var genericObject structs.Generic
var connection *websocket.Conn

var currentBlock structs.Block
var currentBuffer structs.Buffer

func dialer() error {
	URL := "ws://" + host + ":" + port + ""
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(URL, nil)
	connection = conn
	return err
}

func connectToServer() {
	var result structs.Generic
	result.Type = "connecting"
	isConnected = true
	fmt.Println("Connecting to: " + host + ":" + port + ".")
	connection.WriteMessage(websocket.TextMessage, genericMarshall(result))
}

func main() {
	err := dialer()
	if err != nil {
		fmt.Println(err)
		return
	}

	connectToServer()
	go iter()
	for {
		time.Sleep(time.Millisecond * 100)
	}
}

func iter() {
	for isConnected == true {
		time.Sleep(time.Millisecond * 400)
		var incomming structs.Generic
		err := connection.ReadJSON(&incomming)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				isConnected = false
				connection.Close()
				reconnectToServer()
			}
			fmt.Println(err)
			return
		}
		decode(incomming)
	}
}

func reconnectToServer() {
	for isConnected != true {
		println("Reconnect to server...")
		time.Sleep(time.Millisecond * 5000)
	}
}

func decode(incommingMessage structs.Generic) {

	genericObject.Type = ""
	genericObject.JSONString = ""

	if incommingMessage.Type == "connected" {
		genericObject.Type = "connected"
		isConnected = true
		fmt.Println("Connected to server.")
		connection.WriteMessage(websocket.TextMessage, genericMarshall(genericObject))
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
			isMining = true
			//miner.Miner(connection, currentBlock, currentBuffer)
			println("Head changed, starting on block:" + strconv.Itoa(currentBlock.Header.Height) + ".")
		} else {
			isMining = true
			// miner.Miner(connection, currentBlock, currentBuffer)
			println("Start mining on block: " + strconv.Itoa(currentBlock.Header.Height) + ".")
		}
	} else if incommingMessage.Type == "stopWorker" {
		isMining = false
		//miner.StopMining()
		println("Stop mining!")
	} else {
		genericObject.Type = "ping"
		genericObject.JSONString = ""
		connection.WriteMessage(websocket.PingMessage, genericMarshall(genericObject))
	}
}

func genericMarshall(generic structs.Generic) []byte {
	result, _ := json.Marshal(generic)
	return result
}
