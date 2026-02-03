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

func main() {
	interviewOne()
}

// https://youtu.be/d25GqHI_plA?si=4QiG4i03nu7w4zsj

/*

:= new variable decler


* Code segment  *
======================

func interviewOne = (){...}



* Data secment *
============================
var a = 10




* Stack *
=============================
Fist time call init - Skip
then call main

age = 30
age > 18  = true
a := 47





* Heap *
==================================



** Print **
Inside if block: 47
Outside if block: 10




*/
