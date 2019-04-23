package main

import "fmt"

func main()  {
	//变量声明
	//1.声明var 变量名 类型   变量声明之后必须使用
	//2.只是声明变量，默认是0
	//3.在同一个{} 变量名是唯一的
	var a int
	var b int
	a,b = 10,20//变量赋值
	fmt.Println(a)
	fmt.Println(b)
	var c,d int
	fmt.Println(c,d)
	//二变量的初始化  声明变量时  同时赋值
	var e int = 10//初始化   一步到位
	e = 20 //先声明后赋值

	fmt.Println(e)

	var cc float64 = 2
	var sum float64 = cc*cc//等号后面可以跟表达式
	fmt.Println(sum)


}
