package examples

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type SocketsImplementation struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (SocketsImplementation) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			serverMsg := []byte(" will reflect in all clients listening")
			serverMsg = append(msg, serverMsg...)
			// Write message back to browser
			if err = conn.WriteMessage(msgType, serverMsg); err != nil {
				return
			}
		}
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./examples/websockets.html")
	})

	http.ListenAndServe(":80", router)
}
