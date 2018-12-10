/*
Example for how to send and receive tokens from the lexer.
When the lexer find something, it creates a token of type tken. It will choose the type of token found and put that into tkenType, and the text found will be put into the argument.
Then the tken is put on the channel to be received by the parser, and the go struct code will be generated within the switch/case selection.
*/

package main

import (
	"fmt"
	"sync"
)

//tkenType is the type describing a token.
// A <start> start tag will have token start.
// An </start> end tag will have token end.
type tkenType int

const (
	start tkenType = iota
	end
)

type tken struct {
	tkenType        //type of token
	arg      string //the actual text found in the xml while lexing
}

var wg sync.WaitGroup

func main() {

	//make a channel of type tken with a buffer of 2 to send tokens on.
	tkCh := make(chan tken, 2)

	//send two test tokens onto the channel of type tken,
	// and then close the channel.
	tkCh <- tken{start, "aSuperFancyThing"}
	tkCh <- tken{end, ""}
	close(tkCh)

	//Since my we are starting the function as a concurrent go routine with the "go"
	// command, we need to add a waitgroup to wait for it to finnish before we exit
	// the program. The program will wait at "wg.Wait()" until the function is done,
	// then continue and exit.
	wg.Add(1)
	go readTokens(tkCh)
	wg.Wait()

}

func readTokens(tkCh chan tken) {
	//Range over all the values found in the channel.
	//When the last item is read the range command will detect that the channel is closed,
	// and exit the for loop.
	for v := range tkCh {
		switch v.tkenType {
		case start:
			fmt.Printf("type %v struct {\n", v.arg)
		case end:
			fmt.Printf("}\n")
		}
	}

	//We are here done with the function. Tell the waitgroup that we are done, to allow the main
	// function to continue and no longer block/hold the program at the "wg.Wait()" command.
	wg.Done()
}
