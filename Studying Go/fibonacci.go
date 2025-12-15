package main

import "fmt"

func main() {

	const fibonacciPosition int = 20

	for i := 0; i < fibonacciPosition; i++ {
		fmt.Printf("%d\n", fibonacci(i))
	}

}

// Create a Fibonacci calculation function
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
