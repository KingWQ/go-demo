package main

import "fmt"

/*定义接口类*/
type usb interface {
	read() string
	write(string) bool
}

type phone struct {
	Name string
}

func (phone phone) Read() string {
	fmt.Println(phone.Name, "read")
	return "读取到的内容"
}

type computer struct {
	Name string
}

func (computer computer) Write(content string) bool {
	fmt.Println(computer.Name, "写入", content)
	return true
}

func main() {
	phone := phone{"iphone11"}
	content := phone.Read()
	fmt.Println(content)

	computer := computer{"macbook pro"}
	computer.Write("代码")
}
