package main

import (
	"fmt"
	"sync"
)

var num int = 100
var wg3 sync.WaitGroup
var lock sync.Mutex

func main() {
	//testNormal()
	testGoroutine()
	//testLock()
}
func testLock() {
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go add3()
	}
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go sub3()
	}
	wg3.Wait()
	fmt.Println(num)
}
func add3() {
	lock.Lock()
	num++
	wg3.Done()
	lock.Unlock()
}
func sub3() {
	lock.Lock()
	num--
	wg3.Done()
	lock.Unlock()
}

func testGoroutine() {
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go add2()
	}
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go sub2()
	}
	wg3.Wait()
	fmt.Println(num)
}
func add2() {
	num++
	wg3.Done()
}
func sub2() {
	num--
	wg3.Done()
}

func testNormal() {
	for i := 0; i < 100; i++ {
		add1()
	}
	for i := 0; i < 100; i++ {
		sub1()
	}
	fmt.Println(num)
}
func add1() {
	num++
}
func sub1() {
	num--
}
