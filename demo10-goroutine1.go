package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("获取当前GOROOT目录：", runtime.GOROOT())
	fmt.Println("操作系统：", runtime.GOOS)
	fmt.Println("逻辑CPU数量：", runtime.NumCPU())
	//设置最大的可同时使用的CPU核数  取逻辑cpu数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n) //一般在使用之前就将cpu数量设置好 所以最好放在init函数内执行

	go func() {
		fmt.Println("start ...")
		runtime.Goexit()
		fmt.Println("end ...")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("main end ...")
}
