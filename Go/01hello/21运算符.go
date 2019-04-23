package main

import "fmt"

func main() {
	a:=10
	a--
	fmt.Println(a)
	a++
	b:=20
	fmt.Println(a)
	fmt.Println(a==b)
	fmt.Println(a!=b)
	fmt.Println(a&b)
	//取地址
	fmt.Printf("%p\n",&a)
	fmt.Printf("%d\n",*&a)
}
