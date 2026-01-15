package main

import (
	"fmt"

	"example.com/mathlib"
	/* go mod init example.com/first-project */)

func main() {

	runuserinformation()
	checkNumber(10)
	var result = mathlib.Add(4, 4)
	fmt.Println("Addition Result:", result)
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

//https://youtu.be/QHmi0VUz5Sk?si=REsiYOj4s4ufgdAb
