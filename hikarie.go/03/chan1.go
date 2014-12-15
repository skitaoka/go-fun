package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)

	for {
		value, ok := <- c
		if !ok {
			break;
		}
		fmt.Println(value)
	}
}

func foo(c chan <- int) {
	for i := 1; i < 5; i++ {
		c <- i
	}
	close(c)
}

