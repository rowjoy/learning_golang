package main

import (
	"fmt"
)

var a = 10

func interviewOne() {
	age := 30
	if age > 18 {
		a := 47
		fmt.Println("Inside if block:", a)
	}

	fmt.Println("Outside if block:", a)
}

// https://youtu.be/d25GqHI_plA?si=4QiG4i03nu7w4zsj
