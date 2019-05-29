package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Tracer is an interface type that implements one function called Trace.
// The Trace function accepts any type, and any number of that type since
// it takes an open interface{} as it's input, and it is variadic with ...
type Tracer interface {
	//
	Trace(...interface{})
}

// tracer is a struct type with an io.Writer as it's only field.
// Using an io.Writer here means we could give f.ex a file/stdout
// or any type with a Write method as input to tracer.out.
// Here we specify where we want to direct the output.
type tracer struct {
	out io.Writer
}

// Trace :
// For the trace struct to become a type that fullfills the interface
// type Tracer, it need to have a method called Trace with a variadic
// input of type open interface{}.
// The function uses Fprint and direct the output to wherever t.out points
// and prints out whatever is inside 'a...'
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New will return an interface type Tracer. This means we can return any type that
// have method with the signature, trace(...interface{}). trace is such a type.
// The underlying type we return here of this interface type Tracer is a struct tracer.
// We set tracer.out to whatever is given to 'w' of type io.Writer as input,
// and return that.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// nilTracer is also a Tracer since it fullfils the requirement of the Tracer interface
// by having a Trace(...interface) method.
// But nil tracer don't do anything, it got no io.Writer field, and generates no output.
type nilTracer struct{}

// Trace generates no output since the function body is empty.
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}

func main() {
	tracer1 := New(os.Stdout)
	ref1 := reflect.TypeOf(tracer1)
	ref2 := reflect.TypeOf(&tracer1)
	fmt.Printf("The interface type of tracer = %v, and the value type = %v\n", ref2, ref1)
	tracer1.Trace("Some stuff we want to print out from trace1")

	// tracer2 should not print anything.
	tracer2 := Off()
	ref3 := reflect.TypeOf(tracer2)
	ref4 := reflect.TypeOf(&tracer2)
	fmt.Printf("The interface type of tracer = %v, and the value type = %v\n", ref4, ref3)
	tracer2.Trace("Some stuff we want to print out from trace2")
}
