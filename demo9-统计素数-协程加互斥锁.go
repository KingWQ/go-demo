package main

import (
	"fmt"
	"sync"
	"time"
)

var resArr1 []int
var wg5 sync.WaitGroup
var lock1 sync.Mutex

func main() {
	//记录开始时间
	start := time.Now().Unix()

	//并发执行
	for i := 0; i < 10; i++ {
		wg5.Add(1)
		go findPrimeNumbers2(i)
	}
	wg5.Wait()
	//记录结束时间
	end := time.Now().Unix()

	fmt.Println(len(resArr1), end-start)
}

func findPrimeNumbers2(startNum int) {
	for i := startNum * 30000; i <= 30000+startNum*30000; i++ {
		res := findPrimeNumber2(i)
		if res {
			lock1.Lock()
			resArr1 = append(resArr1, i)
			lock1.Unlock()
		}
	}
	wg5.Done()
}

// 判断某个数是否为素数
func findPrimeNumber2(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
