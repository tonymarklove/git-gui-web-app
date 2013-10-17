package main
 
import (
  "code.google.com/p/go.net/websocket"
  "log"
  "net/http"
  "text/template"
)
 
var homeTempl = template.Must(template.ParseFiles("home.html"))
 
func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTempl.Execute(c, req.Host)
}

func main() {
  log.Print("Starting...")

  gitStartup()

  go h.run()

  http.Handle("/", http.FileServer(http.Dir("data")))
  http.Handle("/ws", websocket.Handler(wsHandler))

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
