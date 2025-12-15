package main

import (
	"fmt"
	"time"
)

func main() {
	var stringToPrint string = "Printando uma mensagem..."
	channel := writeStuff(stringToPrint)

	const printAmountOfTimes int = 10

	for i := 0; i < printAmountOfTimes; i++ {
		fmt.Println(<-channel)
	}
}

func writeStuff(textLine string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Inputed value: %s", textLine)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
