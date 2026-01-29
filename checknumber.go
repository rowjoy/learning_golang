package main

import (
	"fmt"
	"strconv"
)

func checkNumber(number int) {
	if number == 10 {
		fmt.Println("Your add value number" + strconv.Itoa(number))
	} else if number%2 == 0 {
		fmt.Println("Your add value numberxx" + strconv.Itoa(number))
	} else {
		fmt.Println("Your not add value number" + strconv.Itoa(number))
	}
}

func runuserinformation() {
	// print welcome message
	fmt.Println("Welcome to the User Information Program")
	var name string
	var age int
	var email string

	// Get user input
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	// Display user information
	fmt.Println("\nUser Information:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Email:", email)
}
