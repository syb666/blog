package main

import "fmt"

//首字母必须为大写才能被其他文件使用
type Student struct {
	id int
	name string
	sex string
	age int
	addr string
}

func main() {
	//初始化顺序，每个成员都得初始化
	var s1 Student
	s1.id = 1
	s1.name = "孙亚博"
	s1.age = 20
	s1.addr = "河南大学"
	fmt.Println(s1)

	//自动推导类型
	s2 := Student{id:2,name:"yano",age:30,addr:"123"}
	fmt.Println(s2)

	//var 变量名  结构体名
	var s3 Student
	fmt.Println(s3)

}
