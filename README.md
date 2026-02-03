# learning_golang

# go code run 
go run filename , other filename 
```example : go run main.go , other.go


#go custom package add 
```go mod init example.com / change any name 

``` Higher Order Function is callback function


```css
/// [First Order Function Vs ]
func addNumber(a, b int) {
	c := a + b
	fmt.Println(c)
}

func totalSum(a int, d int) {
	f := a + d*a
	fmt.Println(f)
}

Higher Order Function
- parameter -> function
- function return
- both 

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

go run main.go => compile it => main => ./main 
go build main.go => compile it => main 

./main


✅ Topics Covered
===========================
Parameter VS Argument
First order function
Init function – you cannot call this; the computer calls this automatically
Standard or named function
Anonymous function
IIFE – Immediately Invoked Function Expression
Function expression or assign function in variable
Higher order function or first class function (treated as first class)
Callback function
Closure – close over

struct & REceiver function --- new 



Variadic function 
Defer function – last in, first out (LIFO)

+---------------------+
|   Code Segment      |
+---------------------+
|   Data Segment      |
+---------------------+
|   Heap (GC)         |
+---------------------+
|   Stack (Frame)     |
+---------------------+

Code Segment (Text Segment)

🔹 এখানে থাকে:

প্রোগ্রামের machine code / instruction
function এর actual code

🔹 Feature:

Read-only
একবার load হলে change হয় না
Data Segment

🔹 এখানে থাকে:

Global variable
Static variable

Data segment দুই ভাগে হয়:

Initialized data
Uninitialized data (BSS)


Heap

🔹 এখানে থাকে:
Runtime এ dynamically allocated memory

🔹 Feature:

Manual / Automatic free
Stack থেকে slow
Size বড়


One Line Summary

Code Segment → Program instruction
Data Segment → Global / static data
Stack → Function + local variable
Heap → Dynamic memory
GC → Heap clean করে



```