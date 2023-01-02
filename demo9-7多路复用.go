package main

import (
	"fmt"
	"strconv"
)

func main() {
	chanInt := make(chan int, 100)
	chanStr := make(chan string, 100)

	counter := 1

	for {
		if counter > 10 {
			break
		}
		chanInt <- counter
		chanStr <- "hello" + strconv.Itoa(counter)

		counter++
	}

	for {
		select {
		case v := <-chanInt:
			fmt.Println(v)
		case v := <-chanStr:
			fmt.Println(v)
		default:
			return
		}
	}
}
