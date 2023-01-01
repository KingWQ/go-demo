package main

import "fmt"

func main() {
	var a int = 100
	fmt.Printf("a的值为：%d\n地址：%p\n", a, &a)

	//使用指针
	var b int = 100
	var p *int = &a
	fmt.Println(p)
	fmt.Println(b)
	*p = 200
	fmt.Println(*p)
	fmt.Println(b)
}
