package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// API handles api homeroute
func API(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world")
}

// ActionSocket is the api action socket
// (listens for garage door commands)
func ActionSocket(w http.ResponseWriter, r *http.Request, app *config.App) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("something went wrong with the websocket: %s", err)
		return
	}
	defer conn.Close()

	log.Printf("Client %s subscribed\n", r.RemoteAddr)
	defer log.Printf("Client %s unsubscribed\n", r.RemoteAddr)
	
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading from websocket:", err)
			break
		} else if mt == websocket.BinaryMessage {
			log.Println("websocket had binary message", err)
			break
		}

		// debug
		log.Printf("websocket received: %s", msg)
	}
}

func internalError(ws *websocket.Conn, msg string, err error) {
	log.Println(msg, err)
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}
