package main

import (
	"fmt"
	"reflect"
)

/*
Interfaces are a tool. Whether you use them or not is up to you,
but they can make code clearer, and they can provide a nice
API between packages, or clients (users) and servers (providers).

Yes, you can create your own struct type,
and you can "attach" methods to it, for example:
*/
type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

/*
There is some similarity in both of the above types: both have
a method Say() with the same signature (parameters and result types).
We can capture this with an interface:
*/

type Sayer interface {
	Say() string
}

/*
The interface contains only the signatures of the methods,
but not their implementation.

Note that in Go a type implicitly implements an interface if its method
set is a superset of the interface. There is no declaration of the intent.
What does this mean? Our previous Cat and Dog types already implement this Sayer interface even though this interface definition didn't even exist when we wrote them earlier, and we didn't touch them to mark them or something. They just do.

Interfaces specify behavior. A type that implements an interface means that
type has all the methods the interface "prescribes".

Since both implement Sayer, we can handle both as a value of Sayer,
they have this in common. See how we can handle both in unity:
*/

type Horse struct{}

func (h Horse) Say() string { return "neigh" }

func MakeCatTalk(c Cat) {
	fmt.Println("Cat says:", c.Say())
}

func MakeTalk(s Sayer) {
	fmt.Println(reflect.TypeOf(s).Name(), "says:", s.Say())
}

func main() {
	c := Cat{}
	fmt.Println("Cat says:", c.Say())
	d := Dog{}
	fmt.Println("Dog says:", d.Say())
	/*
	  We can already see some repetition in the code above: when making
	  both Cat and Dog say something. Can we handle both as the same kind of entity,
	  as animal? Not really. Sure we could handle both as interface{},
	  but if we do so, we can't call their Say() method because a value
	  of type interface{} does not define any methods.
	*/

	animals := []Sayer{c, d}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}

	/*
	  (That reflect part is only to get the type name, don't make much of it as of now.)

	  The important part is that we could handle both Cat and Dog as the
	  same kind (an interface type), and work with them / use them.
	  If you were on quickly to create additional types with a Say() method,
	  they could line up besides Cat and Dog:
	*/

	animals = append(animals, Horse{})
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}

}
