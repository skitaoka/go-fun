package main

import "fmt"

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)

	go func() { for { a <- 0 } }()
	go func() { for { b <- 0 } }()
	go func() { for { c <- 0 } }()
	go func() { for { d <- 0 } }()

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Println("recv a")
		case <-b:
			fmt.Println("recv b")
		case <-c:
			fmt.Println("recv c")
		default:
			fmt.Println("nothing")
		}
	}
}

