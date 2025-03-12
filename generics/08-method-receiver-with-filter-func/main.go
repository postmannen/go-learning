package main

import "fmt"

// Slice is a generic slice that can hold any type
// The type is set when the slice is created.
type Slice[T any] []T

// Filter is a method on the generic Slice type.
// It takes a function as its argument. The function f are the actual filter function that
// can does the filtering logic. This can be acomparison, check for length, check if equal, etc.
func (s Slice[T]) Filter(f func(T) bool) Slice[T] {
	var result Slice[T]

	// Range the slice and run the provided function f for each element.
	for _, v := range s {
		// If the function f returns true, append the element to the result slice.
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	// Run the filter on an int slice, and provide a function that filters for ecen numbers.
	s1 := Slice[int]{1, 2, 3, 4, 5}
	even := s1.Filter(func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(even)

	// Run the filter on a string slice, and provide a function that filters for strings with more than 3 characters.
	s2 := Slice[string]{"grevling", "ku", "hest", "ape", "katt", "hund"}
	filtered := s2.Filter(func(v string) bool {
		return len(v) > 3
	})
	fmt.Println(filtered)
}
