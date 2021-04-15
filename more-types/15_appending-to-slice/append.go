package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
