package main

import "fmt"

func main() {
	testContinue()
	testBreak()
	testGoto()
}

func testContinue() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}

func testBreak() {
	for i := 0; i < 2; i++ {
		for a := 0; a < 5; a++ {
			if a == 2 {
				break
			}
			fmt.Println(i, a)
		}
	}
}

func testGoto() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			goto labelTest
		}
		fmt.Println(i)
	}
labelTest:
	fmt.Println("goto 跳转标签")
}
