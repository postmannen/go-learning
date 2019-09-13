/*
Testing the concept of channels who contains another channel of struct{}
for signaling.
*/
package main

import (
	"fmt"
	"time"
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
		quitCh: make(chan chan struct{}),
	}
	// Start the loop which will check the channel
	go s.loop()

	return s

}

// loop will do select on a time.Tick channel, and also check if we have
// received a quit signal in the form of a chan struct{}
func (m *storage) loop() {
	for {
		select {
		case <-time.Tick(time.Second):
			// time.Tick have a time.Time channel, and we will now receive
			// a value there every second.
			fmt.Println("--- loop method: time...")
		case q := <-m.quitCh:
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

func (m *storage) stop() {
	// to stop the loop method we create a new variable of type chan struct{}
	// which we will put on the quitCh.
	q := make(chan struct{})

	fmt.Println("*** stop method: put 'chan struct{}' on 'm.quitCh'")

	// we put a chan struct{} on the channel,
	m.quitCh <- q

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

func main() {
	m := newstorage()
	time.Sleep(time.Second * 3)
	m.stop()

	fmt.Println("----------------2nd run-------------------------")

	time.Sleep(time.Second * 3)
	m.stop()
}
