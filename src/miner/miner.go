package miner

import (
	"time"

	"../structs"

	"github.com/gorilla/websocket"
)

var connection *websocket.Conn
var IsMining = true

var currentBlock structs.Block
var currentBuffer structs.Buffer

//Miner initialize
func Miner(conn *websocket.Conn, Block structs.Block, Buffer structs.Buffer) {
	connection = conn
	currentBlock = Block
	currentBuffer = Buffer
	_miner()
}

func _miner() {
	for IsMining == true {
		println("MINING")
		time.Sleep(time.Millisecond * 1000)
	}
}

//StopMining function
func StopMining() {
	IsMining = false
}
