package trace

//Tracer is the interface that describes an object capable of
// tracing events throughout code.

//Tracer interface
type Tracer interface {
	Trace(...interface{})
}
