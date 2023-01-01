package main

import "fmt"

func main() {
	//1 指定长度并赋值
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	//2 指定长度后分步赋值
	var arr2 [3]int
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	fmt.Println(arr2)

	//3 自动长度
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	//1 遍历数组 for
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//2 遍历数组 for range
	for _, v := range arr3 {
		fmt.Println(v)
	}

}
