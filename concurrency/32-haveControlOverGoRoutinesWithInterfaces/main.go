package main

import (
	"fmt"
	"sync"
	"time"
)

// processes holds the general structure for controlling the
// go routines started.
type processes struct {
	// register of started and active go routines
	active map[int]struct{}
	wg     sync.WaitGroup
	// Channel to wait for signal from a go routine that it is done.
	allDoneCh chan chan int
	// Channel to send new do'er types to be started as go routines.
	newProcessesCh chan doer
}

// newProcesses will return a structure who holds all the
// information about processes being executed, and will
// also start the checking that all processes have been done.
func newProcesses() *processes {
	ps := processes{
		active:         make(map[int]struct{}),
		allDoneCh:      make(chan chan int),
		newProcessesCh: make(chan doer),
	}

	return &ps
}

// start will kick off the process allowing to register and
// start new processes.
func (ps *processes) start() {
	ps.wg.Add(1)
	ps.checkAllDone()
	ps.execute()
}

// done are called when we are done orchestrating the go routines
// to call, and we want to just wait until they all are finished.
func (ps *processes) done() {
	close(ps.newProcessesCh)
	ps.wg.Wait()
}

// checkAllDone will check the allDoneCh for incomming 'chan int'
// signaled from the go routines when they are done, and we can
// then clean up the register of active go routines.
func (p *processes) checkAllDone() {
	go func() {
		fmt.Printf("Before ranging allDoneCh\n")
		for doneCh := range p.allDoneCh {
			pID := <-doneCh
			fmt.Printf("map content before delete: %#v\n", p.active)
			fmt.Printf("deleting pID: %v\n", pID)
			delete(p.active, pID)

		}
		p.wg.Done()
	}()

}

// Execute the functions 'do' methods on all the 'doer' comming in on the
// newProcesses channel.
func (p *processes) execute() {
	go func() {

		// We need to use a WaitGroup here so we know that all the go
		// routines started are done before we close the allDoneCh.
		var wg sync.WaitGroup
		for v := range p.newProcessesCh {
			p.active[v.getID()] = struct{}{}
			wg.Add(1)
			go v.do(&wg)
		}
		wg.Wait()
		close(p.allDoneCh)
	}()
}

// The control structure of a single process.
type process struct {
	id     int
	doneCh chan chan int
}

type doer interface {
	do(*sync.WaitGroup)
	getID() int
}

// ---

// ProcOne is a single process type to be run.
type procOne struct {
	process process
}

// do is the function who do the actual thing wished
func (p procOne) do(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("procOne %v: %v\n", p.process.id, i)
		time.Sleep(time.Millisecond * 50)
	}

	// Signal back to the checkAllDone() function that this function is done, and
	// can be cleaned up.
	c := make(chan int, 1)
	c <- p.process.id
	p.process.doneCh <- c

	// We need to signal here back to function where this function was called from,
	// since it is ran as a go routine, and the parent function will exit
	// before the <- c is put on the doneCh above, and the parent function will quit
	// to early and close the allDoneCh to early.
	wg.Done()

}

func (p procOne) getID() int {
	return p.process.id
}

func main() {
	ps := newProcesses()
	ps.start()

	// Create a couple of processes, and send them to be
	// scheduled for execution.
	ps.newProcessesCh <- procOne{process: process{
		id:     1,
		doneCh: ps.allDoneCh,
	}}
	time.Sleep(time.Second * 3)
	ps.newProcessesCh <- procOne{process: process{
		id:     2,
		doneCh: ps.allDoneCh,
	}}

	ps.done()
}
