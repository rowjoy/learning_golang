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
