package main

import "fmt"

func main()  {
	//自动推导
	var a = 10
	fmt.Println(a)
	fmt.Printf("%T\n",a)//打印类型
	//常用的自动推导
	//自动推导  :=  左边变量名没使用过
 	b:=20
	fmt.Println(b)

 	c:=3.14
 	fmt.Println(c)
 	fmt.Printf("%T\n",c)//默认float64

 	e,f,g :=10,20,30
 	fmt.Println(e,f,g)
	fmt.Printf("%T-%T-%T",e,f,g)

}
