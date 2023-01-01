package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1 整型和字符串相互转换
	var a int = 10
	var b string = "10"

	fmt.Printf("整型转换为字符串：%s\n", strconv.Itoa(a))
	tmp, _ := strconv.Atoi(b)
	fmt.Printf("字符串转换为整型：%d\n", tmp)

	//2 浮点型和字符串互相转换
	f1 := 0.123456789
	s1 := strconv.FormatFloat(f1, 'f', -1, 32)
	s2 := strconv.FormatFloat(f1, 'f', -1, 64)
	fmt.Println("将浮点数转化为float32转换为string", s1)
	fmt.Println("将float64转换为string", s2)
	s3 := "0.123456789"
	f2, _ := strconv.ParseFloat(s3, 64)
	f3, _ := strconv.ParseFloat(s3, 32)
	fmt.Println("将string转换为float64", f2)
	fmt.Println("将string转换为float32再强制转换为float64", f3)

	//3 整型和浮点型互相转换
	var x int = 10
	var y float32 = 3.14
	fmt.Printf("整型转换为浮点型：%f\n", float32(x))
	fmt.Printf("整型转换为浮点型64：%f\n", float64(x))
	fmt.Printf("浮点型转换为整型：%d\n", int(y))
}
