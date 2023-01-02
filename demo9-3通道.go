package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个整型通道
	//ch := make(chan int)
	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		fmt.Println(i)
	//	}
	//	ch <- 1
	//}()
	//fmt.Println("main done")
	//<-ch

	//chanFor()

	//chanForRange()

	chanDemo()
}

func chanDemo() {
	//创建管道
	channel := make(chan int)

	//开启协程接收管道值，此处会产生阻塞等待管道有数据
	go pOut(channel)

	//发送管道值
	for i := 0; i < 10; i++ {
		channel <- i
	}

	//通知并发协程退出
	channel <- -1

	//主程序结束
	fmt.Println("main end")
}
func pOut(channel chan int) {
	for {
		data := <-channel
		if data == -1 {
			break
		}
		fmt.Println(data)
	}
}

func chanFor() {
	//创建管道
	ch := make(chan int)

	//协程取值
	go func() {
		for {
			data := <-ch
			if data == -1 {
				break
			}
			fmt.Println("取出：", data)
		}
	}()

	//向管道发送数据
	for i := 1; i <= 10; i++ {
		fmt.Println("放入：", i)
		ch <- i
		//每次发送完等待
		time.Sleep(time.Second)
	}

	//通知读取协程退出
	ch <- -1

	fmt.Println("main done")
}

func chanForRange() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("main done")
}
