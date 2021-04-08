// A Message contains a pointer to another message.
// The idea here is to check what happens to that
// pointer the the next message value.
//
// The pointer gets dereferenced when marshaled, so the
// actual content that is being pointed to is put into
// the marshaled data.
// So when we unmarshal the data again it will put in
// the content of what the otherMsg pointed to into
// the new variable.

package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Data     string
	OtherMsg *Message
}

func main() {
	var jsb []byte

	{
		var err error

		m1 := Message{Data: "message-one-data"}
		m2 := Message{Data: "message-two-data", OtherMsg: &m1}

		fmt.Printf("m2: %#v\n", m2)

		jsb, err = json.Marshal(m2)
		if err != nil {
			fmt.Printf("error: json marshal failed: %v\n", err)
		}

		fmt.Printf("jsb1: %s\n", jsb)
	}

	{
		fmt.Printf("jsb2: %s\n", jsb)
		var m3 Message

		err := json.Unmarshal(jsb, &m3)
		if err != nil {
			fmt.Printf("error: json marshal failed: %v\n", err)
		}

		fmt.Printf("m3.OtherMsg: %T, m3.OtherMsg.Data: %s\n", m3.OtherMsg, m3.OtherMsg.Data)
	}

}
