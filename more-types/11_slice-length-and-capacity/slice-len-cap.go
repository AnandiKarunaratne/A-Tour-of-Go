package main

import "fmt"

// A slice has both a length and a capacity.
// length = number of elements it contains.
// capacity = number of elements in the underlying array, counting from the first element in the slice.

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// extend its length
	s = s[:4]
	printSlice(s)

	// drop its first two values
	s = s[2:]
	printSlice(s)

	s = s[:3]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
