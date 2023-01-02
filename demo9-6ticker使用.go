package main

import (
	"fmt"
	"time"
)

func main() {
	//testTicker()
	testTickerChan()
}

func testTickerChan() {
	ch := make(chan int)
	ticker := time.NewTicker(time.Second)

	go func() {
		counter := 0
		for _ = range ticker.C {
			counter++
			ch <- counter
		}
	}()

	for {
		time.Sleep(time.Second)
		select {
		case v := <-ch:
			fmt.Println("tick", v)
			if v >= 10 {
				ticker.Stop()
			}
		default:
			fmt.Println("default")
		}
	}
}

func testTicker() {
	ticker := time.NewTicker(time.Second)
	counter := 0

	for _ = range ticker.C {
		counter++
		fmt.Println(counter)
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}
