var bodyParser = require("body-parser");
const express = require('express');
const http = require('http');
const url = require('url');
const WebSocket = require('ws');
var cors = require('cors')

const app = express();

var isMining = false;

// app.use(cors())
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(cors())
 
app.post('/startWorker', function (req, res) {
  let object = {
    "type": "startWorker",
    "jsonstring": ''+JSON.stringify(req.body.block)+'',
    "jsonstring2": ''+JSON.stringify(req.body.buffer)+''
  }
  sendBroadcast(object)
  if(isMining){
    console.log("Broacast to clients: head-changed.")
  }
  else{
    console.log("Broacast to clients: startWorker.")
    isMining = true
  }
  res.send("Broadcast: start miners.")
})

app.post('/stopWorker', function(req, res){
  let object = {
    "type": "stopWorker",
    "jsonstring": ""
  }
  sendBroadcast(object)
  isMining = false
  console.log("Broacast to clients: stopWorker.")
  res.send("Broadcast: stop miners.")
})
 
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

function sendBroadcast(object) {
      wss.clients.forEach(function each(client) {
        client.send(JSON.stringify(object));
      });
      return
}

wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    message = JSON.parse(message)
    let object = {
                    "type": "",
                    "jsonstring": ""
                  }
    if(message.type == "connecting"){
      console.log("Peer connecting.")
      object.type = "connected"
    }
    else if (message.type == "connected") {
      console.log("Peer connected.")
      object.type = "connecting"
    }
    else if (message.type == "nextPackage") {
      console.log("Sending test data")
      object.type = "nextPackage"
      object.jsonstring = '{"naam": "rick", "leeftijd": 21}'
    }
    else if (message.type == "hashlist"){
      console.log("recieved hashlist")
    }
    else if (message.type == "broadcast"){
    }
    ws.send(JSON.stringify(object))
  });
});


server.listen(8080, function listening() {
  console.log('Listening on %d', server.address().port);
});