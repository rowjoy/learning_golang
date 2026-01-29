package main

import (
	"fmt"
	// "example.com/mathlib"
	/* go mod init example.com/first-project */)

/*
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
	//Function Expression

	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(2, 4)

	add := func(r, t, q int) {
		x := r + t*q
		fmt.Println(x)
	}

	add(2, 4, 6)

}

// First invoked (call) function
func init() {
	fmt.Println("Hello i am is jamirul islam")
}

// https://youtu.be/5vTZqDv6Uis?si=lGBaxxujmGvctWWL -

//https://youtu.be/9iODfk6HMpM?si=SjAJyfg7bJjInDf4
*/

//Parameter Vs Argument | First Order Function Vs Higher Order Function
/*
func add(a, b, c int) { // parameter => a , b , c
	x := a + b + c
	fmt.Println(x)
}

func main() {
	add(2, 3, 4) // argument => 2, 5
}
*/

//https://youtu.be/2hjQO2_zUpk?si=q9oMshlsUVSkRWBT

// /[First Order Function Vs ]
func addNumber(a, b int) {
	c := a + b
	fmt.Println(c)
}

func totalSum(a int, d int) {
	f := a + d*a
	fmt.Println(f)
}

// Higher Order Function
func orocressOpration(x, i int, op func(a, b int)) {
	op(x, i)
}

//  Higher Order Function return function

func call() func(a int, b int) {
	return addNumber
}

/// Higher Order Function is callback function

func main() {
	orocressOpration(3, 4, addNumber)
	orocressOpration(2, 3, totalSum)
	totalNumber := call()
	totalNumber(2, 4)

}
