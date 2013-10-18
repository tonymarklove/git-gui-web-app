package main

import (
  "log"
  "code.google.com/p/go.net/websocket"
)

type connection struct {
  // The websocket connection.
  ws *websocket.Conn
 
  // Buffered channel of outbound messages.
  send chan string
}

var c = connection{}

func (c *connection) reader() {
  for {
    var message string
    err := websocket.Message.Receive(c.ws, &message)

    response := Git(message)
    websocket.Message.Send(c.ws, response)

    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func (c *connection) writer() {
  for message := range c.send {
    err := websocket.Message.Send(c.ws, message)
    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
  log.Print("Connection incoming")

  c.ws = ws
  c.send = make(chan string, 256)

  go c.writer()
  c.reader()
}
