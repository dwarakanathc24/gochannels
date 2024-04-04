package main

import "fmt"

func main() {
	fmt.Println("hello world")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	send(even, odd, quit)

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

func receive(e, o, q <-chan int){
	for{
		select{
		case v:=<-e
		
		}
	}
}
