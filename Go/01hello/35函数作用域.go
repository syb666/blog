package main

import "fmt"
//全局变量名不能和其他文件中的变量名冲突

var a int

func text3()  {
	a:=19
	fmt.Println(a)
}

func main() {
	a = 9
	text3()
	fmt.Println(a)
}
