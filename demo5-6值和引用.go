package main

import "fmt"

type Language struct {
	Name string
	Age  int
}

func main() {
	testNoChange()
	testChange()
}

func testNoChange() {
	var p = Language{
		"Golang",
		12,
	}
	var p2 = p
	p2.Name = "PHP"
	fmt.Printf("%#v 类型 %T\n", p, p)
	fmt.Printf("%#v 类型 %T\n", p2, p2)
}

func testChange() {
	var slice = []int{1, 2, 3, 4, 5}
	slice2 := slice
	slice2[0] = 99
	fmt.Println(slice)
	fmt.Println(slice2)
}
