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

type webData struct {
	divID int
}

func (d *webData) echoHandler(w http.ResponseWriter, r *http.Request) {
	//upgrade the handler to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error: websocket Upgrade: ", err)
	}

	for {
		divStart := fmt.Sprintf("<div id=%v>", d.divID)
		divEnd := fmt.Sprintf("</div>")
		buttonDelete := fmt.Sprintf("<button id=%v>button %v</button>", d.divID, d.divID)

		//read the message
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error: websocket ReadMessage: ", err)
			return
		}

		//print message to console
		fmt.Printf("Client=%v typed : %v \n", conn.RemoteAddr(), string(msg))

		//check if message from client contains a "do action" keyword,
		//if it matches a case, change the content of 'msg', before it is sendt back to the web browser
		strMsg := string(msg)
		switch strMsg {
		case "button1":
			msg = []byte("<button>Test button</button>")
		case "button2":
			m := fmt.Sprint("<script>", "addButtonJS(100)", "</script>")
			msg = []byte(m)
		case "input":
			m := fmt.Sprint(divStart, "<input placeholder='put something here'></input>", buttonDelete, divEnd)
			msg = []byte(m)
			d.divID++
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
