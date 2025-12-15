package main

import "fmt"

func main() {
	const fibonacciPosition int = 40

	tasks := make(chan int, fibonacciPosition)
	results := make(chan int, fibonacciPosition)

	go worker(tasks, results)

	for i := 0; i < fibonacciPosition; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < fibonacciPosition; i++ {
		resultToPrint := <-results
		fmt.Println(resultToPrint)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

// Create a Fibonacci calculation function
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	} else {
		return fibonacci(posicao-2) + fibonacci(posicao-1)
	}
}
