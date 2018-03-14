package main

import (
	"fmt"
)

func appendToSlice1(s []string, test string) {
	s = append(s, "some data1")
	fmt.Println(test, ": appending to slice, and slice contains ", s)
}

func appendToSlice3(s *[]string, test string) {
	*s = append(*s, "some data3")
	fmt.Println(test, ": appending to slice, and slice contains ", s)
}

func appendToSlice4(s []string, test string) []string {
	return append(s, "some data4")
}

func changeExisting(s []string, test string) {
	for i := 0; i < len(s); i++ {
		s[i] = "overwritten data"
	}
}

var mySlice1 []string

func main() {
	fmt.Println("--------------------------")
	//test1
	//This is not working.
	//It is only the header of a slice that is passed to the function,
	//it is also the header who contains the length and capacity of the slice.
	//The function below are able to append to the slice, but since the capacity in the header
	//is not changed, the change is just local in the function, and never passed back to main where the function was called,
	//and the original slice remains unchanged.
	for i := 1; i <= 5; i++ {
		appendToSlice1(mySlice1, "test1")
	}

	fmt.Println("test1 result = ", mySlice1)
	fmt.Println("--------------------------")

	//test2
	//This does not work either, because the append will add to the last elementh (length+1),
	//and since the length is in the header, and the slice header is passed as a value to the function,
	//then the new length made in function will not be reflected back to main.
	mySlice2 := make([]string, 10, 20)
	fmt.Printf("The len of slice = %v, and the cap of slice = %v \n", len(mySlice2), cap(mySlice2))
	for i := 1; i <= 5; i++ {
		appendToSlice1(mySlice2, "test2")
	}

	fmt.Println("test2 result = ", mySlice2)
	fmt.Println("--------------------------")

	//test2.1
	//This works, since we are only using spots in the slice that are defined in the header,
	//and now the changes are also reflected back to the slice in main.
	changeExisting(mySlice2, "test 2.1")

	fmt.Println("test2.1 result = ", mySlice2)
	fmt.Println("--------------------------")

	//test3
	//By using pointers it is no longer a value of the slice header that is passed to the function,
	//but a pointer to the slice itselves, so now we can append new elements to the slice and
	//the changes are changed in the header of the slice in main.
	mySlice3 := []string{}

	for i := 1; i <= 5; i++ {
		appendToSlice3(&mySlice3, "test3")
	}

	fmt.Println("test3 result = ", mySlice3)
	fmt.Println("--------------------------")

	//test4
	//By creating a function who returns the appended slice we can pass back a new slice with a new header that reflects the new lengt,
	//we can then store in our original variable, and have new appended slice.
	//Calling the function here will probably be like wrapping append in a function who does the more or less the same as append to,
	//but it is for testing the different ways of working with slices.
	mySlice4 := []string{}
	for i := 1; i <= 5; i++ {
		mySlice4 = appendToSlice4(mySlice4, "test4")
		fmt.Println("Test4 : mySlice4 now contains : ", mySlice4)
	}

	fmt.Println("test4 result = ", mySlice4)
	fmt.Println("--------------------------")
}
