package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// shorter way to write a sequence of if - else statements
	// It runs the first case whose value is equal to the condition expression.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		// windows etc.
		fmt.Printf("%s. \n", os)
	}
	// break statement is provided automatically.
	// switch cases need not be constants
	// values involved need not be integers.
}
