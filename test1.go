package main

import "fmt"
import "runtime"
import "time"

type person struct {
	name string
	age  int
}

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	//defer, gjør så det du deferer blir utsatt og først kjørt når den funksjonen du er
	//i har kjørt ferdig
	defer fmt.Println("...........og så er programmet slutt")

	fmt.Println("Dette er en test")
	fmt.Println(add(10, 10))

	//for løkke
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
	fmt.Print(swap("Hello ", "World "))
	fmt.Print("\n")

	var ii int
	for ii = 0; ii <= 10; ii++ {
		fmt.Print(ii)
	}
	fmt.Print("\n")

	//for å endre type på en variabel
	var myIntVariable int = 100
	var myFloatVariable float32 = float32(myIntVariable)
	fmt.Println("Verdien av float variabel hentet og konvertert fra int er ", myFloatVariable)

	const truth bool = true
	fmt.Println("Verdien av constant 'truth' = ", truth)

	//for loop
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Println("Summen av for løkka kjørt 10 ganger = ", sum)

	//Go har ikke while, bruk bare for, men dropp init av variabel, og increment. Det vil si det før
	//første semikolon og etter siste semikolon
	sum = 1
	for sum < 1000 {
		sum = sum + sum
		fmt.Print(sum, ",")
	}
	fmt.Println()

	//switch statement
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("else, and it is ", os)
	}

	//time
	today := time.Now().Weekday()
	fmt.Println("Variabelen today er nå = ", today)

	//pointers
	fmt.Println("--------------------")
	iii := 1900
	fmt.Println("iii = ", iii)
	pointer1 := &iii
	fmt.Println("verdien a pointer1 som peker til iii er nå =", *pointer1)
	*pointer1 = 1901
	fmt.Println("Så setter vi verdien til *pointer1 som igjen peker til iii med verdien ", *pointer1)
	fmt.Println("Og iii som fikk ny verdi via *pointer1 er da = ", iii)

	//structs
	fmt.Println("********Structs*********")
	fmt.Println(person{"bob", 10})
	s := person{"jill", 12}
	fmt.Println(s.name, s.age)
	s.age = 14
	fmt.Println(s.name, s.age)
	fmt.Println("Hvis man bare spesifiserer s som skal skriver ut uten .name og .age, da får man = ", s)
	//hvis du lager en peker til en struct så kan du droppe * for pekeren for dereferencig
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//arrays
	//størelsen på ett array settes før typen, det vil si hvis man vil lage ett array av størelsen 10
	//for 'hus' av typen int, så skrives det som : var hus [10]int
}
