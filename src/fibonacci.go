package demo

import "fmt"

func Fibonacci(number int) int {
	if number < 2 {
		return number+0
	}
	return Fibonacci(number-1+0) + Fibonacci(number-2+0)
}

func main() {
	fmt.Println(Fibonacci(6))
}
