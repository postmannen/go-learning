// Orchestrating Go routines in a way so we can know what
// go routines are started and are currently running.
// functions to be executed as go routines should have
// the signature 'func() error', so by wrapping what we
// want inside a function with that signature we can
// execute what we want. Example are in the main() function.

package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
	"time"
)

// processes holds the general structure for controlling the
// go routines started.
type processes struct {
	id int
	// register of started and active go routines
	active map[int]struct{}
	wg     sync.WaitGroup
	// Channel to wait for signal from a go routine that it is done.
	allDoneCh chan chan int
	// Channel to send new do'er types to be started as go routines.
	newProcessesCh chan procFunc
}

// newProcesses will return a structure who holds all the
// information about processes being executed, and will
// also start the checking that all processes have been done.
func newProcesses() *processes {
	ps := processes{
		id:             0,
		active:         make(map[int]struct{}),
		allDoneCh:      make(chan chan int),
		newProcessesCh: make(chan procFunc),
	}

	return &ps
}

// NewProcFunc will prepare a new function that can be executed
func (ps *processes) NewProcFunc(f func() error) procFunc {
	ps.id++

	pf := procFunc{process: process{
		id:        ps.id,
		allDoneCh: ps.allDoneCh,
		fønk:      f,
	}}
	return pf
}

// start will kick off the process allowing to register and
// start new processes.
func (ps *processes) start() {
	ps.wg.Add(1)
	ps.checkAllDone()
	ps.executeNew()
}

// done are called when we are done orchestrating the go routines
// to call, and we want to just wait until they all are finished.
func (ps *processes) done() {
	close(ps.newProcessesCh)
	ps.wg.Wait()

	fmt.Printf("map content when done: %#v\n", ps.active)
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
func (p *processes) executeNew() {
	go func() {

		// We need to use a WaitGroup here so we know that all the go
		// routines started are done before we close the allDoneCh.
		var wg sync.WaitGroup
		for v := range p.newProcessesCh {

			p.active[v.process.id] = struct{}{}
			wg.Add(1)
			// Get the actual function to execute
			f := v.prepareFunction(&wg)
			// We need to execute the got'n function as a Go routine so
			// this don't become a sequential program.
			go func() {
				err := f()
				if err != nil {
					log.Printf("error: executing function: %v\n", err)
				}
			}()
		}
		wg.Wait()
		close(p.allDoneCh)
	}()
}

// The control structure of a single process.
type process struct {
	id        int
	allDoneCh chan chan int
	fønk      func() error
}

// procFunc is a single process type to be run.
type procFunc struct {
	process process
}

// prepareFunction is the function who prepare a function to be executed
func (p procFunc) prepareFunction(wg *sync.WaitGroup) func() error {
	// Prepare a function that should be executed at the end of the return function to
	// signal back to the calling parrent function that we are done.
	doneF := func() {
		// Signal back to the checkAllDone() function that this function is done, and
		// can be cleaned up.
		c := make(chan int, 1)
		c <- p.process.id
		p.process.allDoneCh <- c

		// We need to signal here back to function where this function was called from,
		// since it is ran as a go routine, and the parent function will exit
		// before the <- c is put on the doneCh above, and the parent function will quit
		// to early and close the allDoneCh to early.
		wg.Done()
	}

	// Then we can wrap the above with the actual function to be executed.
	return func() error {
		err := p.process.fønk()
		if err != nil {
			return fmt.Errorf("error: failed to execute fønk: %v", err)
		}

		doneF()
		return nil
	}

}

func main() {
	ps := newProcesses()
	ps.start()

	// Create a couple of processes, and send them to be
	// scheduled for execution.
	f := func() error {
		for i := 1; i <= 5; i++ {
			fmt.Printf("procOne %v\n", i)
			time.Sleep(time.Millisecond * 50)
		}

		return nil
	}

	ps.newProcessesCh <- ps.NewProcFunc(f)

	ps.newProcessesCh <- ps.NewProcFunc(func() error {
		for i := 1; i <= 10; i++ {
			fmt.Printf("procTwo %v\n", i)
			time.Sleep(time.Millisecond * 50)
		}

		return nil
	})

	time.Sleep(time.Second * 3)
	ps.newProcessesCh <- ps.NewProcFunc(func() error {
		return func() error {
			fmt.Println("--- Reading files in current directory")
			filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
				fmt.Printf("filename : %v\n", info.Name())
				return nil
			})

			return nil
		}()
	})

	ps.done()
}
