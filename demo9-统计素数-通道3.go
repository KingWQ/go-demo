package main

import (
	"fmt"
	"sync"
	"time"
)

var wg7 sync.WaitGroup

func main() {
	//记录开始时间
	start := time.Now().Unix()

	//定义存放1-12万的数字通道
	intChan := make(chan int, 1000)
	//存放素数管道
	primeChan := make(chan int, 50000)

	//存放数据
	wg7.Add(1)
	go putNum1(intChan)

	grNum := 2

	//从数字管道中取值判断是否素数并保存到素数管道
	//开启多个协程
	//此处开启了多个协程向素数管道写入数据
	//而打印数据函数需要关闭管道，那么如何关闭管道？
	//当16个协程执行完毕时关闭管道
	//利用一个管道记录协程是否执行完毕
	doneChan := make(chan int, grNum)
	for i := 0; i < grNum; i++ {
		wg7.Add(1)
		go primeNum1(intChan, primeChan, doneChan)

	}

	//打印数据
	//wg2.Add(1)
	//go printNum(primeChan)

	//判断16次协程是否执行完毕
	go func() {
		for i := 0; i < grNum; i++ {
			<-doneChan
		}
		//关闭素数统计管道
		close(primeChan)
		close(doneChan)
	}()

	wg7.Wait()

	fmt.Println(len(primeChan), time.Now().Unix()-start)
}

// 向数字管道内存放数据
func putNum1(intChan chan int) {
	for i := 2; i < 300000; i++ {
		intChan <- i
	}
	close(intChan)
	wg7.Done()
}

// 从数字管道中取值判断是否素数并保存到素数管道
func primeNum1(intChan chan int, primeChan chan int, doneChan chan int) {
	for num := range intChan {
		res := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				res = false
				break
			}
		}
		if res {
			primeChan <- num
		}
	}
	wg7.Done()

	//记录协程执行完毕
	doneChan <- 1
}

// 打印数据
func printNum1(primeChan chan int) {
	//for data := range primeChan {
	//	fmt.Println(data)
	//}
	wg7.Done()
}
