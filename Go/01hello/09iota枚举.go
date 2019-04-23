package main

import "fmt"

func main() {
	//1. iota常量生成器，每一行，自动加1
	//2. iota给常量赋值使用
	const (
		a = iota  //0
		b = iota  //1
		c = iota  //2
	)
	fmt.Println(a,b,c)
	//3. iota遇到const时就会进行重置为0
	const d = iota
	const e = iota
	fmt.Println(d)  //0
	fmt.Println(e)  //0

	//4. 在一个括号里，可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1,b1,c1)  //0,1,2

	//5.如果是同一行,值是一样的
	const(
		j = iota
		j1,j2,j3 = iota,iota,iota
		k = iota
	)
	fmt.Println(j)//0
	fmt.Println(j1,j2,j3)//1,1,1
	fmt.Println(k)//2




}
