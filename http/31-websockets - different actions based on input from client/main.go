package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//upgrade the handler to a websocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("error: websocket Upgrade: ", err)
		}

		for {
			//read the message
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("error: websocket ReadMessage: ", err)
				return
			}

			//print message to console
			fmt.Printf("Client=%v typed : %v \n", conn.RemoteAddr(), string(msg))

			//check if message from client is special, and change the response if it is by chaning the content of msg
			strMsg := string(msg)
			switch strMsg {
			case "button":
				msg = []byte("<button>Test button</button>")
			case "input":
				msg = []byte("<input placeholder='put something here'></input>")
			default:
			}

			//write message back to browser
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				fmt.Println("error: WriteMessage failed :", err)
				return
			}

		}
	}
}

func rootHandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	}
}

func secondHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	}

}

func main() {
	http.HandleFunc("/echo", echoHandler())
	http.HandleFunc("/", rootHandle())
	http.HandleFunc("/second", secondHandle())

	http.ListenAndServe(":8080", nil)

}
