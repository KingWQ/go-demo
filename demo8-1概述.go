package main

import "fmt"

var globalVar int = 10

var initName string

func init() {
	fmt.Println("init run ...")
	initName = "hello world"
}

func main() {
	fmt.Println(globalVar)
	//testVarUse(99)
	type addFunc func(a int, b int) (int, string)
	var addFuncVar addFunc = testAdd
	a, b := addFuncVar(1, 2)
	fmt.Println(a, b)

	res := addDouble(1, 2, testAdd1)
	fmt.Println(res)

	funRun := giveMeAFunc()
	fmt.Println(funRun(1, 2))

	//匿名函数
	testNoName()

	//递归函数
	fmt.Println(testRecursion(0), "should be 100")

	//defer关键字
	testDefer()

	//异常处理
	fmt.Println(testException(10, 0))

	//init函数
	fmt.Println(initName)
}

//异常处理
func testException(a int, b int) float32 {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err: ", err)
		}
	}()

	if b == 0 {
		panic("除数不能为0")
	}

	fmt.Println("函数内部继续运行")
	return float32(a) / float32(b)
}

//defer关键字
func testDefer() int {
	var res int = 0
	defer func() {
		fmt.Println("defer 最后执行")
		fmt.Println("res值为：", res)
	}()

	for i := 0; i < 1000000; i++ {
		res = res + i
	}
	return res
}

//递归函数
func testRecursion(startNumber int) int {
	startNumber += 5
	if startNumber < 100 {
		fmt.Println(startNumber)
		return testRecursion(startNumber)
	} else {
		return startNumber
	}
}

//匿名函数
func testNoName() {
	var a int = 10
	var b int = 20
	smallVal := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}(a, b)
	fmt.Println(smallVal)
}

func testAdd1(a int, b int) int {
	return a + b
}

//函数作为参数
func addDouble(a int, b int, f func(int, int) int) int {
	return f(a, b) * 2
}

//函数作为返回值
func giveMeAFunc() func(int, int) int {
	return testAdd1
}

func testAdd(a int, b int) (int, string) {
	var str string = "测试字符串返回值"
	if a > 10 {
		return a, str
	}
	return a + b, str
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func testVarUse(i int) {
	//d为局部变量
	//i为形式参数
	//作用范围为本行内部
	var d int = i
	fmt.Println(d)

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
