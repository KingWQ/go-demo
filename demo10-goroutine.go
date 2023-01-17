package main

import (
	"fmt"
	"time"
)

func main() {
	go goroutine1()
	go goroutine2()

	for i := 1; i <= 5; i++ {
		fmt.Println("main 函数执行---", i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("main 函数结束")
}

func goroutine1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("goroutine1---", i)
	}
}

func goroutine2() {
	for i := 1; i <= 10; i++ {
		fmt.Println("goroutine2---", i)
	}
}
