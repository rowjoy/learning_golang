package main

import (
	"fmt"
)

const x = 10 // constant

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + x + p
		fmt.Println(money)
	}
	return show
}

func call() {
	intcr1 := outer()
	intcr1()
	intcr1()

	intcr2 := outer()
	intcr2()
	intcr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("=== bank ===")
}

// https://youtu.be/Lbq0YwlyDOY?si=5hNwrL04jGZnGovf

/*

 Go run procress
 =============================================

 2 phases
 - compilation phase // compile time
 - exexution phase // run time



 RAM
 =============================================

 **** Code Segment - store all finction code ********
 [read only file add code segmint section]
  x = 10 // const variable not change
  outer = func() {}
  show =  func() {}






 data Segment  - Store all data variable
 Stack - Stack frame - n time need code exexution
 Heap




 Go code work process
 =================================================
 Terninal : go build closure.go
----
Create closure build binary file
-------
run binary file : ./closure











*/
