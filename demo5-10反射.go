package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {
	//reflect1()
	//reflect2()
	//reflect3()
	reflect4()
}

// 1: 通过反射获取类型对象和值对象
func reflect1() {
	a := 36
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

	b := "hello word"
	bType := reflect.TypeOf(b)
	fmt.Println(bType.Name())
	bValue := reflect.ValueOf(b)
	fmt.Println(bValue.String())
}

// 2: 从类型对象中获取类型名称和种类
func reflect2() {
	typeOfStruct := reflect.TypeOf(myStruct{})
	fmt.Println(typeOfStruct.Name())
	fmt.Println(typeOfStruct.Kind())

	typeOfStruct1 := reflect.TypeOf(&myStruct{})
	//获取指针类型指向的元素类型的名称
	fmt.Println(typeOfStruct1.Elem().Name())
	//获取指针类型指向的元素类型的种类
	fmt.Println(typeOfStruct1.Elem().Kind())
}

// 3：获取结构体成员类型
func reflect3() {
	typeOfStruct := reflect.TypeOf(myStruct{})
	fieldNum := typeOfStruct.NumField()

	for i := 0; i < fieldNum; i++ {
		fieldName := typeOfStruct.Field(i)
		fmt.Println(fieldName)

		name, err := typeOfStruct.FieldByName("Name")
		fmt.Println(name, err)
	}
}

func reflect4() {
	a := "golang"
	obj := reflect.ValueOf(a)
	fmt.Println(obj)
	fmt.Println(obj.Interface())
	fmt.Println(obj.String())

	b := myStruct{
		"梅西",
		1,
		31,
	}

	bObject := reflect.ValueOf(b)
	for i := 0; i < bObject.NumField(); i++ {
		fmt.Println(bObject.Field(i).Interface())
	}
	fmt.Println(bObject.Field(1).Type())

	var c *int
	fmt.Println(reflect.ValueOf(c).IsNil())
	fmt.Println(reflect.ValueOf(c).IsValid())
	fmt.Println(reflect.ValueOf(nil).IsValid())

}
