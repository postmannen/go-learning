package main

type intType int

func (i intType) decodeArg() int {
	return int(i)
}

type float32Type float32

func (f float32Type) decodeArg() float32 {
	return float32(f)
}

// ---------------------------------
type command struct{}

type arg1 struct {
	height intType
}

type command1 command

func (c command1) decodeCommand() interface{} {
	return &arg1{height: 20}
}

// --------

type arg2 struct {
	battery float32
}

type command2 command

func (c command2) decodeCommand() interface{} {
	return &arg2{battery: 200.5}
}

// ------------------------------------

func main() {

}
