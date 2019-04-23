package main

import "fmt"

func main() {
	//byte 字符类型 同时也是uint8类型
	// 1. 单引号代表的是字符  双引号 字符串
	var a byte = 'a'
	fmt.Println(a)  //97
	//%c是一个占位符  表示打印一个字符
	fmt.Printf("%c",a)
	fmt.Printf("%c",97)
	fmt.Printf("%T\n",a)

	var b byte = '0'
	fmt.Println(b)   //48
	// \n  \t    \0 字符串结束的标志
	fmt.Println('\n')
	fmt.Println('\t')
}
