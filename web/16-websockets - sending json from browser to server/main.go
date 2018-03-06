package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type webData struct {
	id int
}

type jsonMsg struct {
	ID       int    `json:"id"`
	Action   string `json:"action"`
	Function string `json:"function"`
	Param1   string `json:"param1"`
}

func (d *webData) echoHandler(w http.ResponseWriter, r *http.Request) {
	//upgrade the handler to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error: websocket Upgrade: ", err)
	}

	var msg jsonMsg
	msg.ID = 100

	for {
		//read the message
		msgType, msgIn, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error: websocket ReadMessage: ", err)
			return
		}

		msg.ID++

		//unmarshal the incomming message, and put it in the struct 'msg'
		//var msg jsonMsg
		err = json.Unmarshal(msgIn, &msg)

		//print message to console
		fmt.Printf("Client=%v typed : %v \n", conn.RemoteAddr(), msg)

		//write message back to browser, and marshal msgOut
		msgOut, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("error: marshaling failed: ", err)
		}
		err = conn.WriteMessage(msgType, []byte(msgOut))
		if err != nil {
			fmt.Println("error: WriteMessage failed :", err)
			return
		}

	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func main() {
	wd := webData{
		id: 0,
	}

	http.HandleFunc("/echo", wd.echoHandler)
	http.HandleFunc("/", rootHandle)

	http.ListenAndServe(":8080", nil)

}
