package main

import "fmt"

func main() {
	a,b,c := 10,20,30
	sum:=a+b+c
	//类型转换  用在不同但相互兼容的类型
	//1.数据类型(变量) int(a)
	//2.数据类型(表达式)
	fmt.Println(sum)
	fmt.Println(float64(sum)/7)

	//var  flag  bool
	//flag = true
	//fmt.Println(int(flag))

	//由低类型转换为高类型   保证数据的精度
	//整型-----浮点型   int8-----int16
	//高类型到低类型，精度会丢失
	var d float32 = 3.1
	var e float64 = 3.2
	fmt.Println(float64(d)+e)



}
