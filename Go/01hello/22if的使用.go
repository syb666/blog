package main

import "fmt"

func main() {
	//变量初始化
	//var score int = 700
	//自动推导
	score := 700
	//if语法结构 条件表达式没有括号
	if score>600 {
		fmt.Println("1")
	}
	//if支持一个初始化语句，初始化语句和判断条件用 ; 分割
	if s:=800;s>700{
		fmt.Println("Yes")
	}




}
