package main

import "time"

func main() {

	myChannel := make(chan string)

	var eg string = "Move in"
	go writeStuff(eg, myChannel)

	print("Move out")
}

func writeStuff(texto string, myChannel chan string) {
	for i := 0; i < 10; i++ {
		myChannel <- texto
		time.Sleep(time.Second)
	}
	close(myChannel)
}
