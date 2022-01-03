package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"unsafe"

	"github.com/fxamacker/cbor/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/vmihailenco/msgpack/v5"
)

type Method string
type Node string

type Message struct {
	_ struct{} `cbor:",toarray"`
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
	mPrev := Message{
		ToNode:          "prevsomenode",
		ToNodes:         nil,
		ID:              1000,
		Data:            []string{"Dec 21 10:13:45 someship env[197267]: 2021/12/21 10:13:45 info: processBufferMessages: done with message, deleting key from bucket, 0"},
		Method:          "someMethodToUse",
		MethodArgs:      []string{"some", "arguments", "here"},
		ReplyMethod:     "someMethodToUse",
		ACKTimeout:      10,
		ReplyACKTimeout: 0,
		MethodTimeout:   10,
		Directory:       "some-directort",
		FileName:        "some-filename",
	}

	m := Message{
		ToNode:          "somenode",
		ToNodes:         nil,
		ID:              1001,
		Data:            []string{"Dec 21 10:13:45 someship env[197267]: 2021/12/21 10:13:45 info: processBufferMessages: done with message, deleting key from bucket, 0"},
		Method:          "someMethodToUse",
		MethodArgs:      []string{"some", "arguments", "here"},
		ReplyMethod:     "someMethodToUse",
		ACKTimeout:      10,
		ReplyACKTimeout: 0,
		MethodTimeout:   10,
		Directory:       "some-directort",
		FileName:        "some-filename",
		PreviousMessage: &mPrev,
	}

	fmt.Printf(" * length of m struct, no serialization : %v\n", unsafe.Sizeof(m))

	fmt.Println(" -------------------------MSGPACK------------------------ ")

	{
		b, err := msgpack.Marshal(m)
		if err != nil {
			panic(err)
		}

		fmt.Printf(" * length of b: %v\n", len(b))

		var out Message
		err = msgpack.Unmarshal(b, &out)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)

		fmt.Printf("PreviousMessage test, id : %v\n", out.PreviousMessage.ID)
	}

	fmt.Println(" ------------------------GOB--------------------------")

	{
		var bufGob bytes.Buffer
		gobEncoder := gob.NewEncoder(&bufGob)
		err := gobEncoder.Encode(m)
		if err != nil {
			log.Printf("error: gob encoding failed: %v\n", err)
			return
		}

		// The second time GOB already have the information about the structure of the data
		// so it no longer takes up much space to send more objects of the same type
		fmt.Printf(" * length of gob encoded data 2 : %v\n", bufGob.Len())

		decoder := gob.NewDecoder(&bufGob)

		var out Message
		err = decoder.Decode(&out)
		if err != nil {
			log.Println("Decode s3:", err)
			return
		}
		log.Println("Decoded value:", out)

		fmt.Printf("PreviousMessage test, id : %v\n", out.PreviousMessage.ID)
	}

	fmt.Println(" ------------------------------JSON-------------------------------")

	{
		mJson, err := json.Marshal(m)
		if err != nil {
			log.Printf("error: json marshal failed: %v\n", err)
			return
		}

		fmt.Printf(" * length of mJson : %v\n", len(mJson))

		var out Message
		err = json.Unmarshal(mJson, &out)
		if err != nil {
			log.Printf("error: failed unmarshaling json: %v\n", err)
			return
		}

		fmt.Printf("PreviousMessage test, id : %v\n", out.PreviousMessage.ID)
	}

	fmt.Println("------------------------------CBOR---------------------------------")

	{
		mCbor, err := cbor.Marshal(m)
		if err != nil {
			log.Printf("error: cbor marshal failed: %v\n", err)
			return
		}

		fmt.Printf(" * length of mCbor : %v\n", len(mCbor))

		var out Message
		cbor.Unmarshal(mCbor, &out)

		if err != nil {
			log.Printf("error: failed unmarshaling cbor: %v\n", err)
			return
		}

		fmt.Printf("PreviousMessage test : %v\n", out.PreviousMessage)

		{
			fmt.Println("---------------zstd SpeedBestCompression compression of cbor----------------------")

			zstdW, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
			if err != nil {
				log.Printf("error: zstd new writer failed: %v\n", err)
				return
			}

			b := zstdW.EncodeAll(mCbor, nil)
			fmt.Printf(" * len of zstd SpeedBestCompression cbor data : %v\n", len(b))
		}

		{
			fmt.Println("---------------zstd std options compression of cbor----------------------")

			zstdW, err := zstd.NewWriter(nil)
			if err != nil {
				log.Printf("error: zstd new writer failed: %v\n", err)
				return
			}

			b := zstdW.EncodeAll(mCbor, nil)
			fmt.Printf(" * len of zstd std compression cbor data : %v\n", len(b))
		}

		fmt.Println("---------------gzip compression of cbor----------------------")
		{
			var bufGzip bytes.Buffer
			gzipW := gzip.NewWriter(&bufGzip)
			n, err := gzipW.Write(mCbor)
			if err != nil {
				log.Printf("error: failed to write gzip: %v\n", err)
			}
			fmt.Printf(" * writer wrote %v number of charachters\n", n)
			gzipW.Close()

			fmt.Printf(" * length of gzipB: %v\n", len(bufGzip.Bytes()))

		}

		// -------------- decode ------------

		//{
		//	zstdR, err := zstd.NewReader(nil)
		//	if err != nil {
		//		log.Printf("error: zstd NewReader failed: %v\n", err)
		//		return
		//	}
		//	out, err := zstdR.DecodeAll(b, nil)
		//	if err != nil {
		//		log.Printf("error: zstd decode failed: %v\n", err)
		//		return
		//	}
		//
		//	var m Message
		//
		//	err = cbor.Unmarshal(out, &m)
		//	if err != nil {
		//		log.Printf("error: cbor unmarshal failed: %v\n", err)
		//		return
		//	}
		//
		//	fmt.Printf("PreviousMessage test cbor : %v\n", m.PreviousMessage)
		//}
	}

}
