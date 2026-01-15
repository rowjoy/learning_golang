package mathlib

// All function name Upper case
func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}
func Divide(a int, b int) int {
	if b == 0 {
		return 0 // handle division by zero
	}
	return a / b
}
func Modulus(a int, b int) int {
	return a % b
}
