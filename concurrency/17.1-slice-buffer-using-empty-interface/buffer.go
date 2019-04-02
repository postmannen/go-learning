/*
 Make a buffered reader of channel.
 Will keep the next input values read in a buffer where size if defined by b.size.
 Will release a new value with the readNext method.
*/
package bufferChanLearning

//buffer is a buffer
type buffer struct {
	ChOut          chan interface{} //ChOut, the channel out to be read by client
	Slice          []interface{}    //Slice which is the actual buffer
	confirmNewRead chan bool        //used to wait for confirmation of grabbing the next value from input channel.
	size           int              //size of buffer
}

//NewBuffer creates a new buffer, takes the buffer size as an input argument, and returns a *buffer.
func NewBuffer(m int) *buffer {
	return &buffer{
		ChOut:          make(chan interface{}),
		confirmNewRead: make(chan bool),
		size:           m,
	}
}

//Start will start filling the buffer, to continue filling buffer  use the readNext method.
func (b *buffer) Start(chIn chan interface{}) {
	go func() {
		for len(b.Slice) < b.size-1 {
			v, ok := <-chIn
			if !ok {
				//log.Println("SERVER: done reading chIn in the slice filling loop at the top")
				break
			}
			b.Slice = append(b.Slice, v)
		}

		//Loop and read another value as long as the slice is > 0.
		// Since we fill the buffer when we start as the first thing
		// the only reason for the length of the slice is 0 is that
		// the input channel is closed, and the decrement of the channel
		// value by value have started.
		for len(b.Slice) > 0 {
			v, ok := <-chIn

			//input channel closed ?
			if !ok {
				b.Slice = b.Slice[1:]
			}

			//length of slice already full ?
			if len(b.Slice) == b.size {
				b.Slice = b.Slice[1:]
				b.Slice = append(b.Slice, v)
			}

			//input channel not closed, and slice buffer not filled to size.
			if ok && len(b.Slice) < b.size {
				b.Slice = append(b.Slice, v)
			}

			//slice have been emptied, break out of for loop.
			if len(b.Slice) == 0 {
				break
			}

			b.ChOut <- b.Slice[0]
			//Wait for confirmation for another read from main.
			<-b.confirmNewRead
		}
		close(b.ChOut)
	}()
}

//ReadNext will, relese the lock on the go routine inside the start method,
//and let it read another value from the incomming channel and put it
//into the buffer.
func (b *buffer) ReadNext() {
	b.confirmNewRead <- true
}
