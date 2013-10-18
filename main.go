package main
 
import (
  "code.google.com/p/go.net/websocket"
  "log"
  "net/http"
)

func main() {
  log.Print("Starting...")

  http.Handle("/", http.FileServer(http.Dir("data")))
  http.Handle("/ws", websocket.Handler(wsHandler))

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
