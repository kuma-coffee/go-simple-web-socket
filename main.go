package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Reader(conn *websocket.Conn) {
	for {
		// read message from client
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// show message
		log.Println(string(p))

		// send message to client
		err = conn.WriteMessage(msgType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	Reader(conn)
}

func main() {
	http.HandleFunc("/ws", WebSocketHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
