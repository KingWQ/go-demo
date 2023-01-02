package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg4 sync.WaitGroup
var mutex sync.RWMutex

func main() {
	var students []string
	//开启协程进行写操作
	wg4.Add(1)
	go func() {
		mutex.RLock()
		for i := 0; i < 100; i++ {
			students = append(students, "姓名"+strconv.Itoa(i))
		}
		time.Sleep(time.Second * 2)
		wg4.Done()
		mutex.RUnlock()
		fmt.Println("写操作完成")
	}()

	//开始协程进行读操作数据
	for i := 0; i < 10; i++ {
		wg4.Add(1)
		go func() {
			fmt.Println("读数据执行中")
			time.Sleep(time.Millisecond * 200)
		}()
		wg4.Done()
	}

	wg4.Wait()

	for _, v := range students {
		fmt.Println(v)
	}

	fmt.Println("main done")
}
