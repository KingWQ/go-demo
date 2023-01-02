package main

import (
	"fmt"
	"time"
)

func main() {
	//记录开始时间
	start := time.Now().Unix()

	var resArr []int
	for i := 2; i <= 300000; i++ {
		res := findPrimeNumber(i)
		if res {
			resArr = append(resArr, i)
		}
	}

	//记录结束时间
	end := time.Now().Unix()

	fmt.Println(len(resArr), end-start)
}

//判断某个数是否为素数
func findPrimeNumber(num int) bool {
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
