package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//1. uint是无符号整型  int有符号整型
	//var a uint  = -10
	//fmt.Println(a)
	b := 20
	fmt.Println(b)
	//%T打印输出类型
	fmt.Printf("%T\n",b)
	fmt.Println(unsafe.Sizeof(b))//在内存中占的字节  8

	var c int32 = 20
	fmt.Printf("%T\n",c)
	fmt.Println(unsafe.Sizeof(c))//在内存中占的字节  4

}
