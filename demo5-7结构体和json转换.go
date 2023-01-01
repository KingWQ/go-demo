package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name string "json:\"name\""
	Age  int    "json:\"age\""
}

type School struct {
	Name     string    "json:\"name\""
	Teachers []Teacher "json:\"teachers\""
}

func main() {
	structToJson()
	jsonToStruct()
	structLabel()
	complicatedStruct()
}

//结构体转换为json字符串
func structToJson() {
	var s = Student{
		"二狗",
		18,
	}
	//函数将结构体转换为byte类型的字符数据
	jsonByte, _ := json.Marshal(s)
	//使用string将字符转换为字符串
	fmt.Println(string(jsonByte))
}

//json字符串转换为结构体
func jsonToStruct() {
	var studentJson = `{"Name":"二狗","Age":18}`
	var student Student
	//将字符串转为结构体 注意利用地址方式赋值给结构体
	err := json.Unmarshal([]byte(studentJson), &student)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%#v\n", student)
	}
}

//结构体标签
func structLabel() {
	var teacherJson = `{"name":"Shirley","age":28}`
	var teacher Teacher
	err := json.Unmarshal([]byte(teacherJson), &teacher)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("%#v\n", teacher)
	}

}

//复杂结构体转换
func complicatedStruct() {
	var schoolFirst School = School{}
	schoolFirst.Name = "第一高中"
	schoolFirst.Teachers = append(schoolFirst.Teachers, Teacher{Name: "张老师", Age: 40})
	schoolFirst.Teachers = append(schoolFirst.Teachers, Teacher{Name: "李老师", Age: 29})
	schoolFirst.Teachers = append(schoolFirst.Teachers, Teacher{Name: "王老师", Age: 33})

	jsonByte, err := json.Marshal(schoolFirst)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonByte))
	}

	var jsonStr = `{"name":"第一高中","teachers":[{"name":"张老师","age":40},{"name":"李老师","age":29},{"name":"王老师","age":33}]}`
	var schoolSecond School = School{}
	err2 := json.Unmarshal([]byte(jsonStr), &schoolSecond)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(schoolSecond.Name)
		for _, teacher := range schoolSecond.Teachers {
			fmt.Println(teacher.Name, teacher.Age)
		}
	}
}
