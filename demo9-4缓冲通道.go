package main

import "fmt"

func main() {
	//testNoBufferChan()
	testBufferChan()
}

func testNoBufferChan() {
	ch := make(chan int)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("管道容量：", len(ch))
	fmt.Println("main done")
}

func testBufferChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("管道容量：", len(ch))
	fmt.Println("main done")
}
