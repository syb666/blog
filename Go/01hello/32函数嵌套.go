package main

import "fmt"

func text01(arg...int) {

	fmt.Println(arg)
	for _,i := range arg{
		fmt.Println(i)
	}
}

func text02(arg...int)  {

	//传递多个参数
	text01(arg[2:3]...)
	text01(arg...)
	//text01(2,3,4,5,6)
}

func main() {
	text02(1,2,3,4,5,6)
}
