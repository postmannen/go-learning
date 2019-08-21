package main

import "fmt"

// Command type
type Command struct {
	Project byte
	Class   byte
	Command byte
}

// Ctor type
type Ctor func() Decoder

// Decoder type
type Decoder interface {
	Decode(buf []byte)
}

func main() {
	message := []byte{1, 2, 3, 4}

	ctors := map[Command]Ctor{}
	ctors[Command{1, 2, 3}] = Message{}.New

	command, data := parse(message)
	if ctor, ok := ctors[command]; ok {
		message := ctor()
		message.Decode(data)
		fmt.Printf("%#v\n", message)
	}
}

// Message type
type Message struct {
	Value byte
}

// New will return a new message
func (Message) New() Decoder { return &Message{} }

// Decode will decode the []byte into m.Value
func (m *Message) Decode(buf []byte) {
	m.Value = buf[0]
}

func parse(data []byte) (Command, []byte) {
	return Command{data[0], data[1], data[2]}, data[3:]
}
