package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	v := <-ch
	fmt.Println(v)
	ch <- 6
}
