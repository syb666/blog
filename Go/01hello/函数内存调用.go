package main

import "fmt"

func add(s1,s2 int) int  {
	sum:=s1+s2
	return sum
}

//内存存储模型
//内存相对于一个可执行文件来说分为四个部分
//代码区  数据区   堆区  栈区

func main() {
	//var a int
	a := add(1,2)
	fmt.Println(a)
}
