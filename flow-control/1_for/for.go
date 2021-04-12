package main

import "fmt"

// Go has only one looping construct, the for loop.
// init, condition, post statements, without paranthesis
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
