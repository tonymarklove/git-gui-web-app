package main
 
import (
  "code.google.com/p/go.net/websocket"
  "log"
  "net/http"
  "time"
)

func main() {
  log.Print("Starting...")

  http.Handle("/", http.FileServer(http.Dir("data")))
  http.Handle("/ws", websocket.Handler(wsHandler))


  ticker := time.NewTicker(5 * time.Second)
  quit := make(chan struct{})
  go func() {
    for {
      select {
        case <- ticker.C:
          changedFiles := GitChangedFiles()
          res := responseMessage{MsgType: "changedFiles", ChangedFiles: changedFiles}
          c.send(res)
        case <- quit:
          ticker.Stop()
          return
      }
    }
  }()

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
