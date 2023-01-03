package main

import (
	"fmt"
	"sync"
	"time"
)

var resArr2 []int
var wg6 sync.WaitGroup
var m = &sync.Mutex{}

func main() {
	//记录开始时间
	start := time.Now().Unix()

	gnum := 4
	num := 300000
	ch := make(chan int, num)
	for i := 1; i <= num; i++ {
		ch <- i
	}
	close(ch)

	//并发执行
	for i := 0; i < gnum; i++ {
		wg6.Add(1)
		go findPrimeNumbers3(ch)
	}
	wg6.Wait()

	//记录结束时间
	end := time.Now().Unix()

	fmt.Println("G:", gnum, "all:", num, "result:", len(resArr2), "time:", end-start)
}

func findPrimeNumbers3(ch chan int) {
	for i := range ch {
		res := findPrimeNumber3(i)
		if res {
			m.Lock()
			resArr2 = append(resArr2, i)
			m.Unlock()
		}
	}
	wg6.Done()
}

// 判断某个数是否为素数
func findPrimeNumber3(num int) bool {
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
