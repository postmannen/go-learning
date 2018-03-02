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
	divID int
}

type jsonMsg struct {
	Function string `json:"function"`
	Param1   string `json:"param1"`
}

func (d *webData) echoHandler(w http.ResponseWriter, r *http.Request) {
	//upgrade the handler to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error: websocket Upgrade: ", err)
	}

	for {
		//read the message
		msgType, msgJson, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error: websocket ReadMessage: ", err)
			return
		}

		var msg jsonMsg
		err = json.Unmarshal(msgJson, &msg)

		//print message to console
		fmt.Printf("Client=%v typed : %v \n", conn.RemoteAddr(), msg)

		//write message back to browser
		err = conn.WriteMessage(msgType, []byte(msg.Function))
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
		divID: 0,
	}

	http.HandleFunc("/echo", wd.echoHandler)
	http.HandleFunc("/", rootHandle)

	http.ListenAndServe(":8080", nil)

}
