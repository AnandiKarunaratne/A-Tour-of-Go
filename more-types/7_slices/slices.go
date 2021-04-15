package main

import "fmt"

// An array has a fixed size.
// A slice is a dynamically-sized, flexible view into the elements of an array.
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// The type []T is a slice with elements of type T.
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// This selects a half-open range which includes the first element, but excludes the last one.
	var s []int = primes[1:4]
	fmt.Println(s)
}
