package main

import "fmt"

func main() {
	//1 声明变量
	var person = map[string]string{
		"name": "Mike",
		"age":  "18",
	}
	person["sex"] = "man"
	fmt.Println(person["name"], person["age"], person["sex"])
	delete(person, "age")
	fmt.Println("---", person["age"], "---")

	//2 make函数
	car := make(map[string]string)
	fmt.Println(car)
	car["brand"] = "audi"
	car["size"] = "A4L"
	fmt.Println(car["brand"], car["size"])
}
