package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}

//socketHandler is the handler who controls all the serverside part
//of the websocket. The other handlers like the rootHandle have to
//load a page containing the JS websocket code to start up the
//communication with the serside websocket.
//This handler is used with all the other handlers if they open a
//websocket on the client side.
func socketHandler() http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
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
	var init sync.Once
	var tpl *template.Template
	var err error

	init.Do(func() {
		tpl, err = template.ParseFiles("websockets1.html")
		if err != nil {
			log.Printf("error: ParseFile : %v\n", err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "websocket", nil)
		//http.ServeFile(w, r, "websockets.html")
	}
}

func secondHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	}

}

func main() {
	http.HandleFunc("/echo", socketHandler())
	http.HandleFunc("/", rootHandle())
	http.HandleFunc("/second", secondHandle())

	http.ListenAndServe(":8080", nil)

}
