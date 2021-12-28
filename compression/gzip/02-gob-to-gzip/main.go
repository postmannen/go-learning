package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"log"
	"unsafe"
)

type Method string
type Node string

type Message struct {
	// The node to send the message to.
	ToNode Node `json:"toNode" yaml:"toNode"`
	// ToNodes to specify several hosts to send message to in the
	// form of an slice/array.
	ToNodes []Node `json:"toNodes,omitempty" yaml:"toNodes,omitempty"`
	// The Unique ID of the message
	ID int `json:"id" yaml:"id"`
	// The actual data in the message. This is typically where we
	// specify the cli commands to execute on a node, and this is
	// also the field where we put the returned data in a reply
	// message.
	Data []string `json:"data" yaml:"data"`
	// Method, what request type to use, like REQCliCommand, REQHttpGet..
	Method Method `json:"method" yaml:"method"`
	// Additional arguments that might be needed when executing the
	// method. Can be f.ex. an ip address if it is a tcp sender, or the
	// shell command to execute in a cli session.
	// TODO:
	MethodArgs []string `json:"methodArgs" yaml:"methodArgs"`
	// ReplyMethod, is the method to use for the reply message.
	// By default the reply method will be set to log to file, but
	// you can override it setting your own here.
	ReplyMethod Method `json:"replyMethod" yaml:"replyMethod"`
	// Additional arguments that might be needed when executing the reply
	// method. Can be f.ex. an ip address if it is a tcp sender, or the
	// shell command to execute in a cli session.
	// TODO:
	ReplyMethodArgs []string `json:"replyMethodArgs" yaml:"replyMethodArgs"`
	// IsReply are used to tell that this is a reply message. By default
	// the system sends the output of a request method back to the node
	// the message originated from. If it is a reply method we want the
	// result of the reply message to be sent to the central server, so
	// we can use this value if set to swap the toNode, and fromNode
	// fields.
	IsReply bool `json:"isReply" yaml:"isReply"`
	// From what node the message originated
	FromNode Node
	// ACKTimeout for waiting for an ack message
	ACKTimeout int `json:"ACKTimeout" yaml:"ACKTimeout"`
	// Resend retries
	Retries int `json:"retries" yaml:"retries"`
	// The ACK timeout of the new message created via a request event.
	ReplyACKTimeout int `json:"replyACKTimeout" yaml:"replyACKTimeout"`
	// The retries of the new message created via a request event.
	ReplyRetries int `json:"replyRetries" yaml:"replyRetries"`
	// Timeout for long a process should be allowed to operate
	MethodTimeout int `json:"methodTimeout" yaml:"methodTimeout"`
	// Timeout for long a process should be allowed to operate
	ReplyMethodTimeout int `json:"replyMethodTimeout" yaml:"replyMethodTimeout"`
	// Directory is a string that can be used to create the
	//directory structure when saving the result of some method.
	// For example "syslog","metrics", or "metrics/mysensor"
	// The type is typically used in the handler of a method.
	Directory string `json:"directory" yaml:"directory"`
	// FileName is used to be able to set a wanted name
	// on a file being saved as the result of data being handled
	// by a method handler.
	FileName string `json:"fileName" yaml:"fileName"`
	// PreviousMessage are used for example if a reply message is
	// generated and we also need a copy of  the details of the the
	// initial request message.
	PreviousMessage *Message

	// The node to relay the message via.
	RelayViaNode Node `json:"relayViaNode" yaml:"relayViaNode"`
	// The original value of the RelayViaNode.
	RelayOriginalViaNode Node `json:"relayOriginalViaNode" yaml:"relayOriginalViaNode"`
	// The node where the relayed message originated, and where we want
	// to send back the end result.
	RelayFromNode Node `json:"relayFromNode" yaml:"relayFromNode"`
	// The original value of the ToNode field of the original message.
	RelayToNode Node `json:"relayToNode" yaml:"relayToNode"`
	// The original method of the message.
	RelayOriginalMethod Method `json:"relayOriginalMethod" yaml:"relayOriginalMethod"`
	// The method to use when the reply of the relayed message came
	// back to where originated from.
	RelayReplyMethod Method `json:"relayReplyMethod" yaml:"relayReplyMethod"`
}

func main() {
	m := Message{
		ToNode:          "somenode",
		ToNodes:         nil,
		ID:              1000,
		Data:            []string{"Dec 21 10:13:45 ww.salome env[197267]: 2021/12/21 10:13:45 info: processBufferMessages: done with message, deleting key from bucket, 0"},
		Method:          "someMethodToUse",
		MethodArgs:      []string{"some", "arguments", "here"},
		ReplyMethod:     "someMethodToUse",
		ACKTimeout:      10,
		ReplyACKTimeout: 0,
		MethodTimeout:   10,
		Directory:       "some-directort",
		FileName:        "some-filename",
	}
	//dataString := `Tue Dec 21 10:22:16 2021 : {ID:10467 Data:{Subject:{ToNode:central CommandOrEvent:EventNACK Method:REQHello messageCh:<nil>} Message:{ToNode:central ToNodes:[] ID:0 D
	//	ata:[Hello from central] Method:REQHello MethodArgs:[] ReplyMethod: ReplyMethodArgs:[] IsReply:false FromNode:central ACKTimeout:10 Retries:1 ReplyACKTimeout:0 ReplyRetries:0 MethodTimeout: 10 ReplyMethodTimeout:0 Directory:hello-messages FileName:hello.log PreviousMessage:<nil> RelayViaNode: RelayOriginalViaNode: RelayFromNode: RelayToNode: RelayOrigina
	//	lMethod: RelayReplyMethod: done:0xc0003ea000}}}
	//
	//	Dec 21 10:13:45 ww.salome env[197267]: 2021/12/21 10:13:45 info: processBufferMessages: done with message, deleting key from bucket, 0`
	//

	fmt.Printf("* length of m : %v\n", unsafe.Sizeof(m))

	var bufGob bytes.Buffer
	gobEncoder := gob.NewEncoder(&bufGob)
	err := gobEncoder.Encode(m)
	if err != nil {
		log.Printf("error: gob encoding failed: %v\n", err)
		return
	}

	// The first time we can see that GOB adds information about the structure for the stream,
	// so the size of the first gob is a lot bigger  than the initial data.
	fmt.Printf("* length of gob encoded data 1 : %v\n", bufGob.Len())

	err = gobEncoder.Encode(m)
	if err != nil {
		log.Printf("error: gob encoding failed: %v\n", err)
		return
	}

	// The second time GOB already have the information about the structure of the data
	// so it no longer takes up much space to send more objects of the same type
	fmt.Printf("* length of gob encoded data 2 : %v\n", bufGob.Len())

	err = gobEncoder.Encode(m)
	if err != nil {
		log.Printf("error: gob encoding failed: %v\n", err)
		return
	}

	fmt.Printf("* length of gob encoded data 3 : %v\n", bufGob.Len())

	var bufGzip bytes.Buffer

	{
		gzipW := gzip.NewWriter(&bufGzip)
		n, err := gzipW.Write(bufGob.Bytes())
		if err != nil {
			log.Printf("error: gzip write failed: %v\n", err)
			return
		}
		log.Printf("info: n=%v gzip bytes written\n", n)
		gzipW.Close()
	}

	fmt.Printf(" * length of compressed data: %v\n", bufGzip.Len())

	//
	//{
	//	zr, err := gzip.NewReader(&bufGzip)
	//	if err != nil {
	//		log.Printf("error: failed to create gzip reader: %v\n", err)
	//	}
	//
	//	b, err := ioutil.ReadAll(zr)
	//	if err != nil {
	//		log.Printf("error: gzip read failed: %v\n", err)
	//		return
	//	}
	//
	//	zr.Close()
	//
	//	fmt.Printf("un-compressed data = %s\n", b)
	//}
}
