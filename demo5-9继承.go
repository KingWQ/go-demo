package main

import "fmt"

type Animal struct {
	Name string
}

func (animal Animal) eat() {
	fmt.Println(animal.Name, "吃")
}
func (animal Animal) sleep() {
	fmt.Println(animal.Name, "睡")
}

type Dog struct {
	Animal
	Age int
}

func main() {
	dog := Dog{
		Animal{"小黑"},
		2,
	}
	dog.eat()
	dog.sleep()
}
