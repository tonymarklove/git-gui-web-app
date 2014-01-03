package main

import (
  "log"
  "fmt"
  "encoding/json"
  "code.google.com/p/go.net/websocket"
)

type connection struct {
  ws *websocket.Conn
}

type commandMessage struct {
  MsgType string `json:"type"`
  Command commandMessageCommand
  File string `json:"file"`
}

type commandMessageCommand struct {
  Command string
}


type responseMessage struct {
  MsgType string `json:"type"`
  Data string `json:"data,omitempty"`
  ChangedFiles []string `json:"changedFiles,omitempty"`
}


var c = connection{}

func (c *connection) send(msg responseMessage) {
  jsonResponse, _ := json.Marshal(msg)
  websocket.Message.Send(c.ws, string(jsonResponse))
}

func (c *connection) reader() {
  for {
    var message []byte
    err := websocket.Message.Receive(c.ws, &message)

    var cmd commandMessage
    json.Unmarshal(message, &cmd)

    if cmd.MsgType == "raw" {
      response := Git(cmd.Command.Command)
      res := responseMessage{MsgType: "raw", Data: response}
      c.send(res)
    } else if cmd.MsgType == "diffFile" {
      response := Git("diff-files", "--unified", cmd.File)
      res := responseMessage{MsgType: "fileDiff", Data: response}
      c.send(res)
    } else if cmd.MsgType == "commit" {
      msg := fmt.Sprintf("'%s'", cmd.File)
      Git("commit", "--all", "--message", msg)
    }

    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
  log.Print("Connection incoming")

  c.ws = ws

  c.reader()
}
