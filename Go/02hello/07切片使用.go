package main

import (
	"fmt"
)

func main() {
	//1.
	var slice [] int

	slice = append(slice,1,2,3)
	fmt.Println(len(slice))//长度
	fmt.Println(cap(slice))//容量

	var  s1 [] int = [] int{1,2,3,4}
	fmt.Println(s1)

	//2.自动推导
	s2:=[] int{1,2,9,3,4}
	fmt.Println(s2)
	//3.
	//长度要小于容量   容量可以省略
	s3:=make([]int,5,10)//长度是5,容量是10
	fmt.Println(s3)
	s3 = append(s3,1,2,4,5,6,7,8)//可以一直添加，容量也可以随着变大
	fmt.Println(s3)

	s4:=make([] int,5)
	for i:=0;i<10;i++{
		s4 = append(s4,i)
	}
	fmt.Println(s4)

}
