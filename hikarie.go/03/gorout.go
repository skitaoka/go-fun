package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(3 * time.Second)
	fmt.Println("fin")
}

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
		time.Sleep(1 * time.Second)
	}
}

