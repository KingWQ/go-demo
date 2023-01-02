package main

import (
	"fmt"
	"time"
)

func main() {
	//开启协程
	go goSay()
	go goSayHi()

	//主线程与协程同时运行，当主线程运行完毕后协程也会被关闭
	//所以主线程应该等待协程运行完成后再退出
	time.Sleep(time.Second * 10)
	fmt.Println("main end")
}

func goSay() {
	for i := 1; i <= 10; i++ {
		fmt.Println("say ...", i)
		time.Sleep(time.Millisecond * 100)
	}
}
func goSayHi() {
	for i := 1; i <= 10; i++ {
		fmt.Println("say hi...", i)
		time.Sleep(time.Millisecond * 100)
	}
}
