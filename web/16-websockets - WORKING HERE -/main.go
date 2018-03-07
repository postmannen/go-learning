/*
1. Planning to send JSON parsed instructions from client -> server about what to do,
2. Then prepare data to send to client (this can be for example next availiable ID for button)
3. Send new JSON parsed intstructions from server -> client about what to do, and what data to fill in
	This can be for example what function to call on the client side in JS, and what ID to use in that function


*/
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type webData struct {
	IndexID int
	Names   []string
}

type jsonMsg struct {
	ID        int    `json:"id"`
	Action    string `json:"action"`
	Parameter string `json:"parameter"`
}

func (d *webData) echoHandler(w http.ResponseWriter, r *http.Request) {
	//upgrade the handler to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error: websocket Upgrade: ", err)
	}

	for {
		//initialize a blank msg of type jsonMsg. Used to compose a message back to client.
		//if you want msg.ID to start at zero after each refresh, move the initialization of 'msg' outside the for loop,
		//and set msg.ID statically to for example one, and have an increment inside the for loop
		var msg jsonMsg
		msg.ID = d.IndexID

		//read the message
		msgType, msgIn, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error: websocket ReadMessage: ", err)
			return
		}

		d.IndexID++

		//unmarshal the incomming message, and put it in the struct 'msg'
		//var msg jsonMsg
		err = json.Unmarshal(msgIn, &msg)
		//print message to console
		fmt.Printf("Client=%v typed : %v \n", conn.RemoteAddr(), msg)

		/*
			msg.Action = "clientaction changed by server"

		}*/

		switch msg.Action {
		case "CLIENTACTION":
			msg.Action = "clientaction changed by server"
		}

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

func (d *webData) rootHandle(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "websockets.html")
	//load template file into 't'
	t, err := template.ParseFiles("websockets.html")
	if err != nil {
		fmt.Println("error: template.ParseFiles: ", err)
	}
	//execute template file and pass in the receiver 'd' of type webData to the template
	t.Execute(w, d)
}

func main() {
	//create some dummy webData to pass in to the Handler functions and templates
	wd := webData{
		IndexID: 0,
		Names:   []string{"Arne", "Knut", "Ole"},
	}

	//HandlerFunctions to call for the given url's
	http.HandleFunc("/echo", wd.echoHandler)
	http.HandleFunc("/", wd.rootHandle)

	//start the web server at port 8080
	http.ListenAndServe(":8080", nil)
}
