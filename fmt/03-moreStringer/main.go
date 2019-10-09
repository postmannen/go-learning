package main

import (
	"fmt"
	"strings"
	"unicode"
)

// text like to have the first letter capitalized.
type text string

//String method satisfies the fmt.Stringer Interface.
//It takes no arguments and returns a string,
//and is used for controlling the format of a variable.
func (t text) String() string {
    // First letter is already capitalized.
	if unicode.IsUpper(rune(t[0])) {
	    return fmt.Sprintf(string(t))
	}
	
	// If first letter is lowercase, capitalize it...
	titled := strings.Title(string(t))
	return titled

}

// An animal like to have it's last letter capitalied as well.
type animal text

func (a animal) String() string{
    lastChr := a[len(a)-1]
    aLastUpper := string(unicode.ToUpper(rune(lastChr)))
    aSlice := string(a[:len(a)-1])
    
    return aSlice+aLastUpper   
}

type fish *text

func main() {
	// Text type got a stringer interface that checks if the first
	// letter is capitalized, and if it's not it will capitalize it.
	a := text("apekatt")
	b := text("Grevling")
	
	fmt.Printf("%v, %v\n",a,b)
	
	// An animal got an underlying type of text, but the animal types
	// stringer interface will check the last chr instead of the first.
	// As we can see of the output, the string method of the text type
	// is not used at all, it is only the string method of the type it
	// is made of that is used, and not some underlying type.
	c := animal("spurv")
    fmt.Println(c)
	
	// Same as above, just testing if it makes a difference in what string
	// methods that will be used...but it is no different.
    d1 := text("kråke")
    d2 := animal(d1)
    fmt.Println(d2)
}
