package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func websocketServer(wsConn *websocket.Conn) { _ = "STUB: not implemented"; return }

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/websocket", websocket.Handler(websocketServer))

	fmt.Println("Open http://localhost:8080 to access this demo")

	panic(http.ListenAndServe(":8080", nil))
}
