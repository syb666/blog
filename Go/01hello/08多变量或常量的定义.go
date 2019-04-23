package main

import "fmt"

func main() {
	//变量    程序可运行期间   可以改变的量  声明关键字var


	//不同变量的定义
	//var a int = 1
	//var b float64 = 2.0
	//var (
	//	a int = 1
	//	b float64  =2.0
	//)
	//自动推导
	//var (
	//	a = 1
	//	b = 2.0
	//)
	a,b := 1,3.22
	fmt.Println(a,b)

	//常量:程序运行期间 不可以更改量  声明关键字  const
	//常量定义
	const (
		c = 10
		d = 4.3
	)
	fmt.Println(c,d)
	const e,f = 4,6.5
	fmt.Println(e,f)
}
