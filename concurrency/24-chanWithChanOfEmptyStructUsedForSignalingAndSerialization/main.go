/*
Testing the concept of channels who contains another channel of struct{}
for signaling.
*/
package main

import (
	"fmt"
)

type storage struct {
	data string
	// actionCh is the actions to perform with the storage.data
	actionCh chan func()
	// quitCh will close down the loop, when a chan struct{} is put on the channel
	quitCh chan chan struct{}
}

func newstorage() *storage {

	s := &storage{
		actionCh: make(chan func()),
		quitCh:   make(chan chan struct{}),
	}
	// Start the loop which will check the channel
	go s.loop()

	return s

}

// loop will do select on a time.Tick channel, and also check if we have
// received a quit signal in the form of a chan struct{}
func (s *storage) loop() {
	for {
		select {
		case f := <-s.actionCh:
			fmt.Println("executing storage function ")
			f()
		case q := <-s.quitCh:
			// We have got the value 'chan struct{}' from the channel, and q now contains
			// another channel of struct{}
			// We then close the channel that was put on the channel
			fmt.Println("--- loop method: received a chan struct{} value on m.quitCh")
			fmt.Println("--- loop method: closing the channel that was received on the m.quitCh")
			close(q)

			fmt.Println("--- loop method: after close of inner channel, before calling return from loop method")
			// And return out of the loop()
			return
		}
	}
}

func (s *storage) stop() {
	// to stop the loop method we create a new variable of type chan struct{}
	// which we will put on the quitCh.
	q := make(chan struct{})

	fmt.Println("*** stop method: put 'chan struct{}' on 'm.quitCh'")

	// we put a chan struct{} on the channel,
	s.quitCh <- q

	fmt.Println("*** stop method: right after put chan struct{} on m.quitCh, then let's wait for  <-q")

	// then it will wait/block here for that inner channel value q to close,
	// and it will continue after the 'close(q)' call in the loop method is called.
	//
	// The advantage of this way and get a  confirmation of closing the inner channel
	// within the channel,  is that we can be sure that the loop go routine is now closed.
	// If it for some reason is not getting closed it will block right here.
	//
	// We also keep the main channel available for a new run later, since we're not closing that one.
	<-q

	fmt.Println("*** stop method: inner chan is now closed, and we received a close of the main quitCh , leaving stop() method, so we can be sure for loop in the loop method is broken, and we've returned from the loop method")
}

func (s *storage) add(text string) {
	s.actionCh <- func() {
		s.data = text
	}
}

func (s *storage) get() string {
	// create a channel for receiving values from the function
	dataC := make(chan string)
	// since the anonymous function below will be executed in the loop go routine
	// there will be no deadlock, and we can use an unbuffered dataC channel.
	s.actionCh <- func() {
		fmt.Println("*** putting data on the dataC channel")
		dataC <- s.data
	}

	fmt.Printf("*** waiting at 'return <-dataC'\n\n")
	return <-dataC
}

func main() {
	s := newstorage()

	s.add("monkey")
	fmt.Println(s.get())

	s.stop()

}
