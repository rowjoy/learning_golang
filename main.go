package main

import (
	"fmt"

	"example.com/mathlib"
	/* go mod init example.com/first-project */)

func main() {

	// runuserinformation()
	// checkNumber(10)
	// interviewOne()
	var result = mathlib.Add(4, 4)
	fmt.Println("Addition Result:", result)

	x := 10 // expression
	fmt.Println(x)

	if x > 0 { // if expression
		fmt.Println("X is greater then 0")
	}

	// anonymous function
	// Immediately Invoked Function Expression
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(2, 4)

}

// First invoked (call) function
func init() {
	fmt.Println("Hello i am is jamirul islam")
}

// https://youtu.be/5vTZqDv6Uis?si=lGBaxxujmGvctWWL -

//https://youtu.be/9iODfk6HMpM?si=SjAJyfg7bJjInDf4
