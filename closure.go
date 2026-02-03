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

func Colsuremain() { // It is main
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
 ==========================================================
 [read only file add code segmint section]
  x = 10 // const variable not change, const value read only
  outer = func() {}
  ounter-anonymous-function-1 =  func() {} // Anonymous function
  call = func() {}
  main = func() {}
  init = func() {} [First time call const value - init - main ]








 ***** ---data Segment  - Store all data variable ****
 =======================================================
 var p = 100






 *** Stack - Stack frame - n time need code exexution ***
 ========================================================
 Stack auto matic clean Up
 -----------------------------
 Stack Frame code exution





 *** Heap *** Heap gerbase data outo clear GC- (Gerbase controller )
 ===========================================

 * Heap sotore use return function code & return function variable data
 * Stack Frame call complete then close stack Frame but same function return data use local varibale then
   Heap store return function code * return function use local variable [not use other function this data]

[ Heap not store real function its store reference number to Code Segment because this function Anonymous function in Outer function ]

 * Some stack Frame check Heap store data
 * Anonymous function store use defult name



 Go code work process
 =================================================
 Terninal : go build closure.go
----
Create closure build binary file
-------
run binary file : ./closure






Heap use case

------------------------
Any function return data then this function close & clear for ram .

But Function reurn function then this function store Heap  RAM section use code secment referance
return function use any local variable this variable store Heap Ram section






*/
