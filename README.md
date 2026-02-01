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




```