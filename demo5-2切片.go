package main

import "fmt"

func main() {
	//testVar()
	//testFor()
	//testCut()
	//testMake()
	//testAppend()
	//testCopy()
	testDelete()
}

func testVar() {
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	var slice2 = []string{1: "apple", 2: "pear", 3: "orange"}
	fmt.Println(slice2)
}

func testFor() {
	var slice = []string{1: "apple", 2: "pear", 3: "orange"}
	for key, val := range slice {
		fmt.Println(key, val)
	}
}

func testCut() {
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice1 = slice[2:5]
	for _, v := range slice1 {
		fmt.Println(v)
	}
}

func testMake() {
	var slice = make([]int, 5, 10)
	fmt.Printf("%d--%d\n", len(slice), cap(slice))
}

func testAppend() {
	var slice = make([]int, 5, 10)
	slice = append(slice, 1, 2, 3, 4, 5, 6)
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("%d--%d\n", len(slice), cap(slice))
}

func testCopy() {
	var slice = make([]int, 5, 10)
	//引用类型的例子
	//var slice2 = slice
	//slice2[0] = 100
	//for _, v := range slice {
	//	fmt.Println(v)
	//}

	var slice3 = make([]int, 5, 10)
	copy(slice3, slice)
	slice3[0] = 100
	for _, v := range slice {
		fmt.Println(v)
	}
}

func testDelete() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = deleteSlice(slice, 3)
	for _, v := range slice {
		fmt.Println(v)
	}

}

func deleteSlice(slice []int, index int) []int {
	fmt.Println(slice[:index])
	fmt.Println(slice[index+1:])
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
	return slice
}
