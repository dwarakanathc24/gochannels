package main

import "fmt"

func main() {
	fmt.Println("hello world")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	receive(even, odd, quit)

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println(" the even numbers are ", v)
		case v := <-o:
			fmt.Println(" the odd numbers are ", v)
		case v := <-q:
			fmt.Println(" the quit numbers are ", v)
			return
		}

	}

}
