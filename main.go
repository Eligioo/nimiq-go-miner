package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type generic struct {
	Type        string `json:"type"`
	JSONString  string `json:"jsonstring"`
	JSONString2 string `json:"jsonstring2"`
}

//Block struct for JSON request
type Block struct {
	Header struct {
		Version  int `json:"_version"`
		PrevHash struct {
			Obj struct {
				Num0  int `json:"0"`
				Num1  int `json:"1"`
				Num2  int `json:"2"`
				Num3  int `json:"3"`
				Num4  int `json:"4"`
				Num5  int `json:"5"`
				Num6  int `json:"6"`
				Num7  int `json:"7"`
				Num8  int `json:"8"`
				Num9  int `json:"9"`
				Num10 int `json:"10"`
				Num11 int `json:"11"`
				Num12 int `json:"12"`
				Num13 int `json:"13"`
				Num14 int `json:"14"`
				Num15 int `json:"15"`
				Num16 int `json:"16"`
				Num17 int `json:"17"`
				Num18 int `json:"18"`
				Num19 int `json:"19"`
				Num20 int `json:"20"`
				Num21 int `json:"21"`
				Num22 int `json:"22"`
				Num23 int `json:"23"`
				Num24 int `json:"24"`
				Num25 int `json:"25"`
				Num26 int `json:"26"`
				Num27 int `json:"27"`
				Num28 int `json:"28"`
				Num29 int `json:"29"`
				Num30 int `json:"30"`
				Num31 int `json:"31"`
			} `json:"_obj"`
		} `json:"_prevHash"`
		BodyHash struct {
			Obj struct {
				Num0  int `json:"0"`
				Num1  int `json:"1"`
				Num2  int `json:"2"`
				Num3  int `json:"3"`
				Num4  int `json:"4"`
				Num5  int `json:"5"`
				Num6  int `json:"6"`
				Num7  int `json:"7"`
				Num8  int `json:"8"`
				Num9  int `json:"9"`
				Num10 int `json:"10"`
				Num11 int `json:"11"`
				Num12 int `json:"12"`
				Num13 int `json:"13"`
				Num14 int `json:"14"`
				Num15 int `json:"15"`
				Num16 int `json:"16"`
				Num17 int `json:"17"`
				Num18 int `json:"18"`
				Num19 int `json:"19"`
				Num20 int `json:"20"`
				Num21 int `json:"21"`
				Num22 int `json:"22"`
				Num23 int `json:"23"`
				Num24 int `json:"24"`
				Num25 int `json:"25"`
				Num26 int `json:"26"`
				Num27 int `json:"27"`
				Num28 int `json:"28"`
				Num29 int `json:"29"`
				Num30 int `json:"30"`
				Num31 int `json:"31"`
			} `json:"_obj"`
		} `json:"_bodyHash"`
		AccountsHash struct {
			Obj struct {
				Num0  int `json:"0"`
				Num1  int `json:"1"`
				Num2  int `json:"2"`
				Num3  int `json:"3"`
				Num4  int `json:"4"`
				Num5  int `json:"5"`
				Num6  int `json:"6"`
				Num7  int `json:"7"`
				Num8  int `json:"8"`
				Num9  int `json:"9"`
				Num10 int `json:"10"`
				Num11 int `json:"11"`
				Num12 int `json:"12"`
				Num13 int `json:"13"`
				Num14 int `json:"14"`
				Num15 int `json:"15"`
				Num16 int `json:"16"`
				Num17 int `json:"17"`
				Num18 int `json:"18"`
				Num19 int `json:"19"`
				Num20 int `json:"20"`
				Num21 int `json:"21"`
				Num22 int `json:"22"`
				Num23 int `json:"23"`
				Num24 int `json:"24"`
				Num25 int `json:"25"`
				Num26 int `json:"26"`
				Num27 int `json:"27"`
				Num28 int `json:"28"`
				Num29 int `json:"29"`
				Num30 int `json:"30"`
				Num31 int `json:"31"`
				View  struct {
				} `json:"_view"`
				ReadPos  int `json:"_readPos"`
				WritePos int `json:"_writePos"`
			} `json:"_obj"`
		} `json:"_accountsHash"`
		NBits     int `json:"_nBits"`
		Height    int `json:"_height"`
		Timestamp int `json:"_timestamp"`
		Nonce     int `json:"_nonce"`
	} `json:"_header"`
	Body struct {
		MinerAddr struct {
			Obj struct {
				Num0  int `json:"0"`
				Num1  int `json:"1"`
				Num2  int `json:"2"`
				Num3  int `json:"3"`
				Num4  int `json:"4"`
				Num5  int `json:"5"`
				Num6  int `json:"6"`
				Num7  int `json:"7"`
				Num8  int `json:"8"`
				Num9  int `json:"9"`
				Num10 int `json:"10"`
				Num11 int `json:"11"`
				Num12 int `json:"12"`
				Num13 int `json:"13"`
				Num14 int `json:"14"`
				Num15 int `json:"15"`
				Num16 int `json:"16"`
				Num17 int `json:"17"`
				Num18 int `json:"18"`
				Num19 int `json:"19"`
			} `json:"_obj"`
		} `json:"_minerAddr"`
		Transactions []interface{} `json:"_transactions"`
	} `json:"_body"`
}

//Buffer struct for JSON request
type Buffer struct {
	Num0   int `json:"0"`
	Num1   int `json:"1"`
	Num2   int `json:"2"`
	Num3   int `json:"3"`
	Num4   int `json:"4"`
	Num5   int `json:"5"`
	Num6   int `json:"6"`
	Num7   int `json:"7"`
	Num8   int `json:"8"`
	Num9   int `json:"9"`
	Num10  int `json:"10"`
	Num11  int `json:"11"`
	Num12  int `json:"12"`
	Num13  int `json:"13"`
	Num14  int `json:"14"`
	Num15  int `json:"15"`
	Num16  int `json:"16"`
	Num17  int `json:"17"`
	Num18  int `json:"18"`
	Num19  int `json:"19"`
	Num20  int `json:"20"`
	Num21  int `json:"21"`
	Num22  int `json:"22"`
	Num23  int `json:"23"`
	Num24  int `json:"24"`
	Num25  int `json:"25"`
	Num26  int `json:"26"`
	Num27  int `json:"27"`
	Num28  int `json:"28"`
	Num29  int `json:"29"`
	Num30  int `json:"30"`
	Num31  int `json:"31"`
	Num32  int `json:"32"`
	Num33  int `json:"33"`
	Num34  int `json:"34"`
	Num35  int `json:"35"`
	Num36  int `json:"36"`
	Num37  int `json:"37"`
	Num38  int `json:"38"`
	Num39  int `json:"39"`
	Num40  int `json:"40"`
	Num41  int `json:"41"`
	Num42  int `json:"42"`
	Num43  int `json:"43"`
	Num44  int `json:"44"`
	Num45  int `json:"45"`
	Num46  int `json:"46"`
	Num47  int `json:"47"`
	Num48  int `json:"48"`
	Num49  int `json:"49"`
	Num50  int `json:"50"`
	Num51  int `json:"51"`
	Num52  int `json:"52"`
	Num53  int `json:"53"`
	Num54  int `json:"54"`
	Num55  int `json:"55"`
	Num56  int `json:"56"`
	Num57  int `json:"57"`
	Num58  int `json:"58"`
	Num59  int `json:"59"`
	Num60  int `json:"60"`
	Num61  int `json:"61"`
	Num62  int `json:"62"`
	Num63  int `json:"63"`
	Num64  int `json:"64"`
	Num65  int `json:"65"`
	Num66  int `json:"66"`
	Num67  int `json:"67"`
	Num68  int `json:"68"`
	Num69  int `json:"69"`
	Num70  int `json:"70"`
	Num71  int `json:"71"`
	Num72  int `json:"72"`
	Num73  int `json:"73"`
	Num74  int `json:"74"`
	Num75  int `json:"75"`
	Num76  int `json:"76"`
	Num77  int `json:"77"`
	Num78  int `json:"78"`
	Num79  int `json:"79"`
	Num80  int `json:"80"`
	Num81  int `json:"81"`
	Num82  int `json:"82"`
	Num83  int `json:"83"`
	Num84  int `json:"84"`
	Num85  int `json:"85"`
	Num86  int `json:"86"`
	Num87  int `json:"87"`
	Num88  int `json:"88"`
	Num89  int `json:"89"`
	Num90  int `json:"90"`
	Num91  int `json:"91"`
	Num92  int `json:"92"`
	Num93  int `json:"93"`
	Num94  int `json:"94"`
	Num95  int `json:"95"`
	Num96  int `json:"96"`
	Num97  int `json:"97"`
	Num98  int `json:"98"`
	Num99  int `json:"99"`
	Num100 int `json:"100"`
	Num101 int `json:"101"`
	Num102 int `json:"102"`
	Num103 int `json:"103"`
	Num104 int `json:"104"`
	Num105 int `json:"105"`
	Num106 int `json:"106"`
	Num107 int `json:"107"`
	Num108 int `json:"108"`
	Num109 int `json:"109"`
	Num110 int `json:"110"`
	Num111 int `json:"111"`
	Num112 int `json:"112"`
	Num113 int `json:"113"`
	Num114 int `json:"114"`
	Num115 int `json:"115"`
	Num116 int `json:"116"`
	Num117 int `json:"117"`
	View   struct {
	} `json:"_view"`
	ReadPos  int `json:"_readPos"`
	WritePos int `json:"_writePos"`
}

var host = "localhost"
var port = "8080"
var isMining = false
var genericObject generic

var currentBlock Block
var currentBuffer Buffer

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
