package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println((err))
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server start")
	})
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe("localhost:8000", nil)
}
