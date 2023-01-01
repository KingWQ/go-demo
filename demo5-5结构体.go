package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() {
	fmt.Println("Hello, i'm", p.Name)
}

func (p Person) SayAge(agePlus int) int {
	fmt.Println("My age is ", p.Age+agePlus)
	return p.Age + agePlus
}

func main() {
	//1 普通方式
	var p1 Person
	p1.Name = "Golang"
	p1.Age = 14
	fmt.Printf("%#v\n", p1)
	fmt.Println("-----------------------")

	//2 指针方式
	var p2 = new(Person)
	p2.Name = "PHP"
	(*p2).Age = 20
	fmt.Printf("%#v\n类型%T\n", p2, p2)
	fmt.Println("-----------------------")

	//3 键值对方式
	var p3 = Person{
		Name: "Java",
		Age:  20,
	}
	//3.1 结构体指针
	fmt.Printf("%#v\n类型%T\n", p3, p3)
	var p4 = &Person{
		Name: "Vue",
		Age:  10,
	}
	fmt.Printf("%#v\n类型%T\n", p4, p4)
	fmt.Println("-----------------------")

	//4 顺序赋值
	var p5 = &Person{
		"Javascript",
		25,
	}
	fmt.Printf("%#v\n类型%T\n", p5, p5)
	fmt.Println("-----------------------")

	//5 结构体函数
	var p6 = Person{
		"Peter",
		28,
	}
	p6.Say()
	p6.SayAge(2)
}
