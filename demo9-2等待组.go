package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//开启协程
	wg.Add(2)
	go wgSay()
	go wgSayHi()

	//主线程与协程同时运行，当主线程运行完毕后协程也会被关闭
	//所以主线程应该等待协程运行完成后再退出
	wg.Wait()

	fmt.Println("执行完毕")
}

func wgSay() {
	for i := 1; i <= 10; i++ {
		fmt.Println("say...", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
func wgSayHi() {
	for i := 1; i <= 10; i++ {
		fmt.Println("say hi...", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
