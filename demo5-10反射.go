package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Id   int
	Name string
}

func (cat Cat) Say() {
	fmt.Printf("cat.Name: %v\n", cat.Name)
}

func main() {
	//初始化结构体
	cat := Cat{
		Id:   1,
		Name: "花花",
	}
	cat.Say()

	//reflect.ValueOf 反射值
	items := reflect.ValueOf(&cat).Elem()
	//利用属性名称获取值
	item := items.FieldByName("Name")
	fmt.Printf("%T, %v\n", items, item)

	//利用kind获取值
	if item.Kind() == reflect.Invalid {
		fmt.Println("不存在的属性")
	}

	//遍历结构体属性
	typeOfST := items.Type()
	for i := 0; i < items.NumField(); i++ {
		itemInFor := items.Field(i)
		fmt.Printf("%v : %v\n", typeOfST.Field(i).Name, itemInFor.Interface())
	}
	//遍历结构体并调用方法
	//注意func(p *...){} 指针模式无法获取方法
	for n := 0; n < items.NumMethod(); n++ {
		fmt.Printf("typeOfST.Method(n).Name: %v\n", typeOfST.Method(n).Name)
		items.Method(n).Call(nil)
	}
}
