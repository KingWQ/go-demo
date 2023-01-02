package main

import (
	"fmt"
	"sync"
	"time"
)

var resArr []int
var wg1 sync.WaitGroup

func main() {
	//记录开始时间
	start := time.Now().Unix()

	//并发执行
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go findPrimeNumbers(i)
	}
	wg1.Wait()
	//记录结束时间
	end := time.Now().Unix()

	fmt.Println(len(resArr), end-start)
}

func findPrimeNumbers(startNum int) {
	for i := startNum * 30000; i <= 30000+startNum*30000; i++ {
		res := findPrimeNumber1(i)
		if res {
			resArr = append(resArr, i)
		}
	}
	wg1.Done()
}

//判断某个数是否为素数
func findPrimeNumber1(num int) bool {
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
